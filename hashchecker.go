package main

import (
    "crypto/sha1"
    "fmt"
    "io"
    "net"
    "os"
    "path/filepath"
    "strings"
)

func getFileHash(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    hash := sha1.New()
    if _, err := io.Copy(hash, file); err != nil {
        return "", err
    }

    return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func writeToFile(path, data, filePath string) error {
    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    if _, err := file.WriteString(path + ":" + data + "\n"); err != nil {
        return err
    }

    return nil
}

func queryHashRegistry(hash string) (string, error) {
    ipList, err := net.LookupIP(fmt.Sprintf("%s.hash.cymru.com", hash))
    if err != nil {
        if strings.Contains(err.Error(), "no such host") {
            return "", nil
        } else {
            return "", err
        }
    }

    for _, ip := range ipList {
        if ip.String() == "127.0.0.2" {
            return hash, nil
        }
    }

    return "", nil
}

func main() {
    dir := "/Users/pswapneel/Downloads/malware" 
    hashesFile := "/Users/pswapneel/Downloads/malware/hashes.txt"

    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() {
            hash, err := getFileHash(path)
            if err != nil {
                return err
            }

            err = writeToFile(path, hash, hashesFile)
            if err != nil {
                return err
            }

            badHash, err := queryHashRegistry(hash)
            if err != nil {
                return err
            }

            if badHash != "" {
                fmt.Println("Bad hash detected:", badHash)
            }
        }

        return nil
    })

    if err != nil {
        fmt.Println("Error while processing files:", err)
    }
}
