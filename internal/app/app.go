package app

import (
    "bufio"
    "context"
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"

    "github.com/AlecAivazis/survey/v2"
    "github.com/fatih/color"
    "github.com/nekowawolf/MSC/internal/chain"
    "github.com/nekowawolf/MSC/internal/config"
    "github.com/nekowawolf/MSC/internal/ui"
    "github.com/nekowawolf/MSC/internal/wallet"
)

type Recipient struct {
    Address string
}

func Run() {
    ui.ShowBanner()

    chains, err := config.LoadAllChains("chains")
    if err != nil {
        color.Red("Error loading chains: %v", err)
        return
    }

    if len(chains) == 0 {
        color.Red("No chain configurations found in 'chains' directory.")
        return
    }

    var chainNames []string
    for _, c := range chains {
        chainNames = append(chainNames, c.Name)
    }

    var selectedChainName string
    prompt := &survey.Select{
        Message: "Select blockchain network:",
        Options: chainNames,
    }
    err = survey.AskOne(prompt, &selectedChainName)
    if err != nil {
        return
    }

    var selectedChain config.ChainConfig
    for _, c := range chains {
        if c.Name == selectedChainName {
            selectedChain = c
            break
        }
    }

    fmt.Printf("> [MSC] Network selected : %s\n", color.CyanString(selectedChain.Name))

    w, err := wallet.LoadWallet()
    if err != nil {
        color.Red("> [MSC] Error loading wallet: %v", err)
        color.Yellow("Please ensure .env file exists and PRIVATE_KEY is set.")
        return
    }
    
    maskedAddr := w.Address[:6] + "..." + w.Address[len(w.Address)-4:]
    fmt.Printf("> [MSC] Wallet loaded    : %s\n", color.GreenString(maskedAddr))

    client, err := chain.NewChainClient(selectedChain, w)
    if err != nil {
        color.Red("> [MSC] Failed to connect to network: %v", err)
        return
    }
    defer client.Close()

    balance, err := client.GetBalance(context.Background())
    if err != nil {
        color.Red("> [MSC] Failed to fetch balance: %v", err)
    } else {
        fmt.Printf("> [MSC] Balance          : %s %s\n", color.YellowString(fmt.Sprintf("%.4f", balance)), selectedChain.Symbol)
    }
    
    fmt.Println(color.WhiteString("------------------------------------------------------"))

    var filePath string
    filePrompt := &survey.Input{
        Message: "Recipients file — make sure the addresses are correct (ENTER to continue)",
        Default: "recipients.txt",
    }
    err = survey.AskOne(filePrompt, &filePath)
    if err != nil {
        return
    }

    recipients, err := loadRecipients(filePath)
    if err != nil {
        color.Red("> [MSC] Error reading file: %v", err)
        return
    }
    
    fmt.Printf("> [MSC] Recipients loaded: %s addresses\n", color.GreenString(fmt.Sprintf("%d", len(recipients))))

    var amountStr string
    amountPrompt := &survey.Input{
        Message: fmt.Sprintf("Amount to send to EACH address (in %s):", selectedChain.Symbol),
    }
    err = survey.AskOne(amountPrompt, &amountStr)
    if err != nil {
        return
    }

    amount, err := strconv.ParseFloat(amountStr, 64)
    if err != nil {
        color.Red("Invalid amount")
        return
    }

    totalAmount := amount * float64(len(recipients))
    fmt.Println(color.WhiteString("------------------------------------------------------"))
    fmt.Printf("Summary:\n")
    fmt.Printf("Network: %s\n", selectedChain.Name)
    fmt.Printf("Total Recipients: %d\n", len(recipients))
    fmt.Printf("Amount per Recipient: %f %s\n", amount, selectedChain.Symbol)
    fmt.Printf("Total Amount to Send: %f %s\n", totalAmount, selectedChain.Symbol)
    fmt.Printf("Delay between TXs: 5 seconds\n")
    
    confirm := false
    confirmPrompt := &survey.Confirm{
        Message: "Do you want to proceed?",
    }
    survey.AskOne(confirmPrompt, &confirm)

    if !confirm {
        fmt.Println("Operation cancelled.")
        return
    }

    fmt.Println(color.WhiteString("------------------------------------------------------"))
    
    for i, r := range recipients {
        current := i + 1
        fmt.Printf("[%d/%d] Sending %f %s → %s ", current, len(recipients), amount, selectedChain.Symbol, r)
        
        ctx := context.Background()
        txHash, err := client.SendTransaction(ctx, r, amount)
        if err != nil {
            color.Red("✘ Failed")
            fmt.Printf("  Error: %v\n", err)
        } else {
            color.Green("✔ Success")
            fmt.Printf("  TX: %s%s\n", selectedChain.Explorer, txHash)
        }

        if current < len(recipients) {
            fmt.Println("Waiting 5 seconds...")
            time.Sleep(5 * time.Second)
        }
    }
    
    fmt.Println(color.GreenString("\nAll transactions completed."))
}

func loadRecipients(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var recipients []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" && strings.HasPrefix(line, "0x") && len(line) == 42 {
            recipients = append(recipients, line)
        }
    }
    return recipients, scanner.Err()
}