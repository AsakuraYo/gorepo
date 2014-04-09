package main

import (
    "crypto/tls"
    "fmt"
    "log"
    "net/smtp"
    "os"
    "strings"
)

const (
    smtpServerAddr = "smtp.revenco.com"
    smtpPort = 465
    userAccount = "liangjianqiang@revenco.com"
    password = "Yuki@070"
)

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Usage:", os.Args[0], "<toEmails> <subject> <body>")
        fmt.Println("Example:", os.Args[0], "\"abc@revenco.com;123@revenco.com\" \"hello\" \"Hello world\"")
        os.Exit(1)
    }
    toEmail     := os.Args[1]
    mailSubject := os.Args[2]
    mailBody    := os.Args[3]

    toEmails := strings.Split(toEmail, ";")

    message := ""
    message += "From: Liang Jianqiang<" + userAccount + ">\r\n"
    message += "To: " + toEmail + "\r\n"
    message += "Subject: " + mailSubject + "\r\n"
    message += "Content-Type: text/plain; charset=GBK\r\n\r\n"
    message += mailBody

    if err := SendMailUsingTLS(smtpServerAddr, smtpPort, userAccount, password, toEmails, []byte(message)); err != nil {
        log.Println(err)
    } else {
        log.Println("E-mail has been sent successfully.");
    }
}

func SendMailUsingTLS(smtpServerAddr string, smtpPort int, userAccount, password string, toEmails []string, msg []byte) (err error) {
    auth := smtp.PlainAuth("", userAccount, password, smtpServerAddr)

    //create smtp client
    var conf tls.Config
    conf.InsecureSkipVerify = true  // accept insecure server
    conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", smtpServerAddr, smtpPort), &conf)
    if err != nil {
        log.Println("Dial error:", err)
        return err
    }
    c, err := smtp.NewClient(conn, smtpServerAddr)
    if err != nil {
        log.Println("NewClient error:", err)
        return err
    }
    defer c.Close()

    if auth != nil {
        if ok, _ := c.Extension("AUTH"); ok {
            if err = c.Auth(auth); err != nil {
                log.Println("Auth error:", err)
                return err
            }
        }
    }

    if err = c.Mail(userAccount); err != nil {
        return err
    }

    for _, addr := range toEmails {
        if err = c.Rcpt(addr); err != nil {
            return err
        }
    }

    w, err := c.Data()
    if err != nil {
        return err
    }

    if _, err = w.Write(msg); err != nil {
        return err
    }

    if err = w.Close(); err != nil {
        return err
    }

    return c.Quit()
}
