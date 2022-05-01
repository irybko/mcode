package mcode

import (
    "log"
    "golang.org/x/crypto/bcrypt"
)
//
func MakePassword(phrase string) ([]byte,string) {
    // Use GenerateFromPassword to hash & salt pwd
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost. 
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, err := bcrypt.GenerateFromPassword([]byte(phrase), bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    // GenerateFromPassword returns a byte slice so we need to convert the bytes to a string and return it
    return hash, string(hash)
}
//
func GetPasswordFromCmdLine() ([]byte, string) {
    // We will use this to store the users input
    var phrase string
    // Read the users input
    Input("Enter a password",&phrase)
    // Use GenerateFromPassword to hash & salt pwd
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost. 
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    return MakePassword(phrase)
}
//
func CheckPassword_1(hashedPwd string, plainPwd []byte) bool {    // Since we'll be getting the hashed password from the DB it
    // will be a string so we'll need to convert it to a byte slice
    byteHash := []byte(hashedPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        log.Println(err)
        return false
    }
    return true
}
//
func CheckPassword_2(passwd1 []byte,passwd2 []byte) bool {
    err := bcrypt.CompareHashAndPassword(passwd1, passwd2)
    if err != nil {
        log.Println(err)
        return false
    }
    return true
}

