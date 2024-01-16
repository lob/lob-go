# Code Snippets

## Address Api


### Retrieve
```bash
curl https://api.lob.com/v1/addresses/adr_fa85158b26c3eb7c \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

fetchedAddress, _, err := apiClient.AddressesApi.Get(context,"adr_fa85158b26c3eb7c").Execute()

if err != nil {
    return err
}
```






### Delete
```bash
curl -X DELETE "https://api.lob.com/v1/addresses/adr_43769b47aed248c2" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

deletedAddress, _, err := apiClient.AddressesApi.Delete(context, "adr_43769b47aed248c2").Execute()

if err != nil {
    return err
}
```



### Create
```bash
curl https://api.lob.com/v1/addresses \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Harry - Office" \
  -d "name=Harry Zhang" \
  -d "company=Lob" \
  -d "email=harry@lob.com" \
  -d "phone=5555555555" \
  -d "address_line1=2261 Market Street" \
  -d "address_line2=Ste 5668" \
  -d "address_city=San Francisco" \
  -d "address_state=CA" \
  -d "address_zip=94114" \
  -d "address_country=US" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)


var addressCreate = *lob.NewAddressEditable()
addressCreate.SetDescription("Harry - Office")
addressCreate.SetName("Harry Zhang")
addressCreate.SetCompany("Lob")
addressCreate.SetEmail("harry@lob.com")
addressCreate.SetPhone("5555555555")
addressCreate.SetAddressLine1("2261 Market Street")
addressCreate.SetAddressLine2("Ste 5668")
addressCreate.SetAddressCity("San Francisco")
addressCreate.SetAddressState("CA")
addressCreate.SetAddressZip("94114")
addressCreate.SetAddressCountry("US")



createdaddress, _, err := apiClient.AddressesApi.Create(context).AddressEditable(addressCreate).Execute()

if err != nil {
    return err
}
```







### List
```bash
curl -X GET "https://api.lob.com/v1/addresses?limit=2" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)
AddressList = apiClient.AddressesApi.List(context).Execute()
if err != nil {
    return err
}
```



































## Postcards Api


### Retrieve
```bash
curl https://api.lob.com/v1/postcards/psc_5c002b86ce47537a \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

fetchedPostcard, _, err := apiClient.PostcardsApi.Get(context,"psc_5c002b86ce47537a").Execute()

if err != nil {
    return err
}
```






### Delete
```bash
curl -X DELETE "https://api.lob.com/v1/postcards/psc_5c002b86ce47537a" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

deletedPostcard, _, err := apiClient.PostcardsApi.Delete(context, "psc_5c002b86ce47537a").Execute()

if err != nil {
    return err
}
```



### Create
```bash
curl https://api.lob.com/v1/postcards \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Demo Postcard job" \
  -d "from=adr_210a8d4b0b76d77b" \
  --data-urlencode "front=<html style='padding: 1in; font-size: 50;'>Front HTML for {{name}}</html>" \
  --data-urlencode "back=<html style='padding: 1in; font-size: 20;'>Back HTML for {{name}}</html>" \
  -d "to[name]=Harry Zhang" \
  -d "to[address_line1]=2261 Market Street" \
  -d "to[address_line2]=Ste 5668" \
  -d "to[address_city]=San Francisco" \
  -d "to[address_state]=CA" \
  -d "to[address_zip]=94114" \
  -d "merge_variables[name]=Harry" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var to = *lob.NewAddressEditable()
to.SetAddressLine1("2261 Market Street")
to.SetAddressLine2("Ste 5668")
to.SetAddressCity("San Francisco")
to.SetAddressState("CA")
to.SetAddressZip("94114")
to.SetAddressCountry(lob.COUNTRYEXTENDED_US)
to.SetDescription("")
to.SetName("Harry Zhang")

var postcardCreate = *lob.NewPostcardEditable()
postcardCreate.SetDescription("Demo Postcard job")
postcardCreate.SetFrom("adr_210a8d4b0b76d77b")
postcardCreate.SetFront("<html style='padding: 1in; font-size: 50;'>Front HTML for {{name}}</html>")
postcardCreate.SetBack("<html style='padding: 1in; font-size: 20;'>Back HTML for {{name}}</html>")
postcardCreate.SetTo(to)


createdpostcard, _, err := apiClient.PostcardsApi.Create(context).PostcardEditable(postcardCreate).Execute()

if err != nil {
    return err
}
```







### List
```bash
curl -X GET "https://api.lob.com/v1/postcards?limit=2" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)
PostcardList = apiClient.PostcardsApi.List(context).Execute()
if err != nil {
    return err
}
```

## SelfMailers Api


### Retrieve
```bash
curl https://api.lob.com/v1/self_mailers/sfm_8ffbe811dea49dcf \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

fetchedSelfMailer, _, err := apiClient.SelfMailersApi.Get(context,"sfm_8ffbe811dea49dcf").Execute()

if err != nil {
    return err
}
```






### Delete
```bash
curl -X DELETE "https://api.lob.com/v1/self_mailers/sfm_8ffbe811dea49dcf" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

deletedSelfMailer, _, err := apiClient.SelfMailersApi.Delete(context, "sfm_8ffbe811dea49dcf").Execute()

if err != nil {
    return err
}
```



### Create
```bash
curl https://api.lob.com/v1/self_mailers \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Demo Self Mailer job" \
  -d "from=adr_210a8d4b0b76d77b" \
  --data-urlencode "inside=<html style='padding: 1in; font-size: 50;'>Inside HTML for {{name}}</html>" \
  --data-urlencode "outside=<html style='padding: 1in; font-size: 20;'>Outside HTML for {{name}}</html>" \
  -d "to[name]=Harry Zhang" \
  -d "to[address_line1]=2261 Market Street" \
  -d "to[address_line2]=Ste 5668" \
  -d "to[address_city]=San Francisco" \
  -d "to[address_state]=CA" \
  -d "to[address_zip]=94114" \
  -d "merge_variables[name]=Harry" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var to = *lob.NewAddressEditable()
to.SetAddressLine1("2261 Market Street")
to.SetAddressLine2("Ste 5668")
to.SetAddressCity("San Francisco")
to.SetAddressState("CA")
to.SetAddressZip("94114")
to.SetAddressCountry(lob.COUNTRYEXTENDED_US)
to.SetDescription("")
to.SetName("Harry Zhang")

var selfMailerCreate = *lob.NewSelfMailerEditable()
selfMailerCreate.SetDescription("Demo Self Mailer job")
selfMailerCreate.SetFrom("adr_210a8d4b0b76d77b")
selfMailerCreate.SetInside("<html style='padding: 1in; font-size: 50;'>Inside HTML for {{name}}</html>")
selfMailerCreate.SetOutside("<html style='padding: 1in; font-size: 20;'>Outside HTML for {{name}}</html>")
selfMailerCreate.SetTo(to)


createdselfMailer, _, err := apiClient.SelfMailersApi.Create(context).SelfMailerEditable(selfMailerCreate).Execute()

if err != nil {
    return err
}
```







### List
```bash
curl -X GET "https://api.lob.com/v1/self_mailers?limit=2" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)
SelfMailerList = apiClient.SelfMailersApi.List(context).Execute()
if err != nil {
    return err
}
```

## Letters Api


### Retrieve
```bash
curl https://api.lob.com/v1/letters/ltr_4868c3b754655f90 \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

fetchedLetter, _, err := apiClient.LettersApi.Get(context,"ltr_4868c3b754655f90").Execute()

if err != nil {
    return err
}
```






### Delete
```bash
curl -X DELETE "https://api.lob.com/v1/letters/ltr_4868c3b754655f90" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

deletedLetter, _, err := apiClient.LettersApi.Cancel(context, "ltr_4868c3b754655f90").Execute()

if err != nil {
    return err
}
```



### Create
```bash
curl https://api.lob.com/v1/letters \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Demo Letter" \
  -d "from=adr_210a8d4b0b76d77b" \
  --data-urlencode "file=<html style='padding-top: 3in; margin: .5in;'>HTML Letter for {{name}}</html>" \
  -d "color=true" \
  -d "to[name]=Harry Zhang" \
  -d "to[address_line1]=2261 Market Street" \
  -d "to[address_line2]=Ste 5668" \
  -d "to[address_city]=San Francisco" \
  -d "to[address_state]=CA" \
  -d "to[address_zip]=94114" \
  -d "merge_variables[name]=Harry" \
  -d "cards[]=card_c51ae96f5cebf3e"
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var to = *lob.NewAddressEditable()
to.SetAddressLine1("2261 Market Street")
to.SetAddressLine2("Ste 5668")
to.SetAddressCity("San Francisco")
to.SetAddressState("CA")
to.SetAddressZip("94114")
to.SetAddressCountry(lob.COUNTRYEXTENDED_US)
to.SetDescription("")
to.SetName("Harry Zhang")

var letterCreate = *lob.NewLetterEditable()
letterCreate.SetDescription("Demo Letter")
letterCreate.SetFrom("adr_210a8d4b0b76d77b")
letterCreate.SetFile("<html style='padding-top: 3in; margin: .5in;'>HTML Letter for {{name}}</html>")
letterCreate.SetColor("true")
letterCreate.SetTo(to)


createdletter, _, err := apiClient.LettersApi.Create(context).LetterEditable(letterCreate).Execute()

if err != nil {
    return err
}
```







### List
```bash
curl -X GET "https://api.lob.com/v1/letters?limit=2" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)
LetterList = apiClient.LettersApi.List(context).Execute()
if err != nil {
    return err
}
```

## Checks Api


### Retrieve
```bash
curl https://api.lob.com/v1/checks/chk_534f10783683daa0 \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

fetchedCheck, _, err := apiClient.ChecksApi.Get(context,"chk_534f10783683daa0").Execute()

if err != nil {
    return err
}
```






### Delete
```bash
curl -X DELETE "https://api.lob.com/v1/checks/chk_534f10783683daa0" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

deletedCheck, _, err := apiClient.ChecksApi.Cancel(context, "chk_534f10783683daa0").Execute()

if err != nil {
    return err
}
```



### Create
```bash
curl https://api.lob.com/v1/checks \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Demo Check" \
  -d "bank_account=bank_8cad8df5354d33f" \
  -d "amount=22.5" \
  -d "memo=rent" \
  --data-urlencode "logo=https://s3-us-west-2.amazonaws.com/public.lob.com/assets/check_logo.png" \
  --data-urlencode "check_bottom=<h1 style='padding-top:4in;'>Demo Check for {{name}}</h1>" \
  -d "from=adr_210a8d4b0b76d77b" \
  -d "to[name]=Harry Zhang" \
  -d "to[address_line1]=2261 Market Street" \
  -d "to[address_line2]=Ste 5668" \
  -d "to[address_city]=San Francisco" \
  -d "to[address_state]=CA" \
  -d "to[address_zip]=94114" \
  -d "merge_variables[name]=Harry" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var to = *lob.NewAddressEditable()
to.SetAddressLine1("2261 Market Street")
to.SetAddressLine2("Ste 5668")
to.SetAddressCity("San Francisco")
to.SetAddressState("CA")
to.SetAddressZip("94114")
to.SetAddressCountry(lob.COUNTRYEXTENDED_US)
to.SetDescription("")
to.SetName("Harry Zhang")

var checkCreate = *lob.NewCheckEditable()
checkCreate.SetDescription("Demo Check")
checkCreate.SetBankAccount("bank_8cad8df5354d33f")
checkCreate.SetAmount(22.5)
checkCreate.SetMemo("rent")
checkCreate.SetLogo("https://s3-us-west-2.amazonaws.com/public.lob.com/assets/check_logo.png")
checkCreate.SetCheckBottom("<h1 style='padding-top:4in;'>Demo Check for {{name}}</h1>")
checkCreate.SetFrom("adr_210a8d4b0b76d77b")
checkCreate.SetTo(to)


createdcheck, _, err := apiClient.ChecksApi.Create(context).CheckEditable(checkCreate).Execute()

if err != nil {
    return err
}
```







### List
```bash
curl -X GET "https://api.lob.com/v1/checks?limit=2" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)
CheckList = apiClient.ChecksApi.List(context).Execute()
if err != nil {
    return err
}
```

## BankAccounts Api


### Retrieve
```bash
curl https://api.lob.com/v1/bank_accounts/bank_8cad8df5354d33f \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

fetchedBankAccount, _, err := apiClient.BankAccountsApi.Get(context,"bank_8cad8df5354d33f").Execute()

if err != nil {
    return err
}
```






### Delete
```bash
curl -X DELETE "https://api.lob.com/v1/bank_accounts/bank_3e64d9904356b20" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

deletedBankAccount, _, err := apiClient.BankAccountsApi.Delete(context, "bank_3e64d9904356b20").Execute()

if err != nil {
    return err
}
```






### List
```bash
curl -X GET "https://api.lob.com/v1/bank_accounts?limit=2" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)
BankAccountList = apiClient.BankAccountsApi.List(context).Execute()
if err != nil {
    return err
}
```





### Verify
```bash
curl https://api.lob.com/v1/bank_accounts/bank_dfceb4a2a05b57e/verify \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \ 
  -d "amounts[]=25" \ 
  -d "amounts[]=63" \ 
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var bankAccountWritable = *lob.NewBankAccountWritable("322271627", "123456789", lob.BANKTYPEENUM_INDIVIDUAL, "Sinead Connor")

createdBankAccount, _, _ := apiClient.BankAccountsApi.Create(context).BankAccountWritable(bankAccountWritable).Execute()

verifyAmounts := []int32{11, 35}
verify := *lob.NewBankAccountVerify(verifyAmounts)

verifiedAccount, _, err := apiClient.BankAccountsApi.Verify(context, createdBankAccount.Id).BankAccountVerify(verify).Execute()

if err != nil {
    return err
}
```

### Create
```bash
curl https://api.lob.com/v1/bank_accounts \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Test Bank Account" \
  -d "routing_number=322271627" \
  -d "account_number=123456789" \
  -d "signatory=John Doe" \
  -d "account_type=company" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)


var bankAccountCreate = *lob.NewBankAccountWritable()
bankAccountCreate.SetDescription("Test Bank Account")
bankAccountCreate.SetRoutingNumber("322271627")
bankAccountCreate.SetAccountNumber("123456789")
bankAccountCreate.SetSignatory("John Doe")
bankAccountCreate.SetAccountType("company")



createdbankAccount, _, err := apiClient.BankAccountsApi.Create(context).BankAccountWritable(bankAccountCreate).Execute()

if err != nil {
    return err
}
```




## Templates Api


### Retrieve
```bash
curl https://api.lob.com/v1/templates/tmpl_c94e83ca2cd5121 \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

fetchedTemplate, _, err := apiClient.TemplatesApi.Get(context,"tmpl_c94e83ca2cd5121").Execute()

if err != nil {
    return err
}
```






### Delete
```bash
curl -X DELETE "https://api.lob.com/v1/templates/tmpl_df934eeda694203" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

deletedTemplate, _, err := apiClient.TemplatesApi.Delete(context, "tmpl_df934eeda694203").Execute()

if err != nil {
    return err
}
```






### List
```bash
curl -X GET "https://api.lob.com/v1/templates?limit=2" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)
TemplateList = apiClient.TemplatesApi.List(context).Execute()
if err != nil {
    return err
}
```





### Update
```bash
curl https://api.lob.com/v1/templates/tmpl_c94e83ca2cd5121 \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Updated description" \
  -d "published_version=vrsn_362184d96d9b0c9"
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var templateWritable = *lob.NewTemplateWritable("<html>Updated HTML for Template 1/html>")
createdTemplate, _, _ := apiClient.TemplatesApi.Create(context).TemplateWritable(templateWritable).Execute()

templateWritable.SetDescription("Template updated")
resp, _, err := apiClient.TemplatesApi.Update(context, createdTemplate.Id).TemplateWritable(templateWritable).Execute()

if err != nil {
    return err
}
```

### Create
```bash
curl https://api.lob.com/v1/templates \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Test Template" \
  --data-urlencode "html=<html>HTML for {{name}}</html>" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)


var templateCreate = *lob.NewTemplateWritable()
templateCreate.SetDescription("Test Template")
templateCreate.SetHtml("<html>HTML for {{name}}</html>")



createdtemplate, _, err := apiClient.TemplatesApi.Create(context).TemplateWritable(templateCreate).Execute()

if err != nil {
    return err
}
```




## TemplateVersions Api


### Retrieve
```bash
curl https://api.lob.com/v1/templates/tmpl_c94e83ca2cd5121/versions/vrsn_534e339882d2282 \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

fetchedTemplateVersion, _, err := apiClient.TemplateVersionsApi.Get(context,"tmpl_c94e83ca2cd5121", "vrsn_534e339882d2282").Execute()

if err != nil {
    return err
}
```



### Update
```bash
curl https://api.lob.com/v1/templates/tmpl_c94e83ca2cd5121/versions/vrsn_534e339882d2282 \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Updated description"
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var templateWritable = *lob.NewTemplateWritable("<html>Updated HTML for Template 1/html>")
var templateVersionWritable = *lob.NewTemplateVersionWritable("<html>Updated HTML for Template 1/html>")
createdTemplate, _, _ := apiClient.TemplatesApi.Create(context).TemplateWritable(templateWritable).Execute()

createdTemplateVersion, _, err := apiClient.TemplateVersionsApi.Create(context, createdTemplate.Id).TemplateVersionWritable(templateVersionWritable).Execute()
createdTemplateVersion.SetDescription("New Template Version")
resp, _, err := apiClient.TemplatesApi.Update(context, createdTemplate.Id, createdTemplateVersion.Id).TemplateVersionWritable(templateVersionWritable).Execute()

if err != nil {
    return err
}
```



### Delete
```bash
curl -X DELETE "https://api.lob.com/v1/templates/tmpl_4aa14648113e45b/versions/vrsn_534e339882d2282" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

deletedTemplateVersion, _, err := apiClient.TemplateVersionsApi.Delete(context, "tmpl_4aa14648113e45b", "vrsn_534e339882d2282").Execute()

if err != nil {
    return err
}
```


### Update
```bash
curl https://api.lob.com/v1/templates/tmpl_4aa14648113e45b/versions/vrsn_534e339882d2282 \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Updated description"
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var templateWritable = *lob.NewTemplateWritable("<html>Updated HTML for Template 1/html>")
var templateVersionWritable = *lob.NewTemplateVersionWritable("<html>Updated HTML for Template 1/html>")
createdTemplate, _, _ := apiClient.TemplatesApi.Create(context).TemplateWritable(templateWritable).Execute()

createdTemplateVersion, _, err := apiClient.TemplateVersionsApi.Create(context, createdTemplate.Id).TemplateVersionWritable(templateVersionWritable).Execute()
createdTemplateVersion.SetDescription("New Template Version")
resp, _, err := apiClient.TemplatesApi.Update(context, createdTemplate.Id, createdTemplateVersion.Id).TemplateVersionWritable(templateVersionWritable).Execute()

if err != nil {
    return err
}
```




### List
```bash
curl -X GET "https://api.lob.com/v1/templates/tmpl_dadaaf7b76c9f25/versions?limit=2" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)
TemplateVersionList = apiClient.TemplateVersionsApi.List(context, "tmpl_4aa14648113e45b").Execute()
if err != nil {
    return err
}
```
### Update
```bash
curl https://api.lob.com/v1/templates/tmpl_dadaaf7b76c9f25/versions/ \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Updated description"
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var templateWritable = *lob.NewTemplateWritable("<html>Updated HTML for Template 1/html>")
var templateVersionWritable = *lob.NewTemplateVersionWritable("<html>Updated HTML for Template 1/html>")
createdTemplate, _, _ := apiClient.TemplatesApi.Create(context).TemplateWritable(templateWritable).Execute()

createdTemplateVersion, _, err := apiClient.TemplateVersionsApi.Create(context, createdTemplate.Id).TemplateVersionWritable(templateVersionWritable).Execute()
createdTemplateVersion.SetDescription("New Template Version")
resp, _, err := apiClient.TemplatesApi.Update(context, createdTemplate.Id, createdTemplateVersion.Id).TemplateVersionWritable(templateVersionWritable).Execute()

if err != nil {
    return err
}
```





### Update
```bash
curl https://api.lob.com/v1/templates/tmpl_c94e83ca2cd5121/versions/vrsn_534e339882d2282 \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Updated description"
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var templateWritable = *lob.NewTemplateWritable("<html>Updated HTML for Template 1/html>")
var templateVersionWritable = *lob.NewTemplateVersionWritable("<html>Updated HTML for Template 1/html>")
createdTemplate, _, _ := apiClient.TemplatesApi.Create(context).TemplateWritable(templateWritable).Execute()

createdTemplateVersion, _, err := apiClient.TemplateVersionsApi.Create(context, createdTemplate.Id).TemplateVersionWritable(templateVersionWritable).Execute()
createdTemplateVersion.SetDescription("New Template Version")
resp, _, err := apiClient.TemplatesApi.Update(context, createdTemplate.Id, createdTemplateVersion.Id).TemplateVersionWritable(templateVersionWritable).Execute()

if err != nil {
    return err
}
```

### Create
```bash
curl https://api.lob.com/v1/templates/tmpl_4aa14648113e45b/versions \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Second Version" \
  --data-urlencode "html=<html>Second HTML for {{name}}</html>" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)


var templateVersionCreate = *lob.NewTemplateVersionWritable()
templateVersionCreate.SetDescription("Second Version")
templateVersionCreate.SetHtml("<html>Second HTML for {{name}}</html>")



createdtemplateVersion, _, err := apiClient.TemplateVersionsApi.Create(context, "tmpl_4aa14648113e45b").TemplateVersionWritable(templateVersionCreate).Execute()

if err != nil {
    return err
}
```



### Update
```bash
curl https://api.lob.com/v1/templates//versions/ \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Updated description"
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var templateWritable = *lob.NewTemplateWritable("<html>Updated HTML for Template 1/html>")
var templateVersionWritable = *lob.NewTemplateVersionWritable("<html>Updated HTML for Template 1/html>")
createdTemplate, _, _ := apiClient.TemplatesApi.Create(context).TemplateWritable(templateWritable).Execute()

createdTemplateVersion, _, err := apiClient.TemplateVersionsApi.Create(context, createdTemplate.Id).TemplateVersionWritable(templateVersionWritable).Execute()
createdTemplateVersion.SetDescription("New Template Version")
resp, _, err := apiClient.TemplatesApi.Update(context, createdTemplate.Id, createdTemplateVersion.Id).TemplateVersionWritable(templateVersionWritable).Execute()

if err != nil {
    return err
}
```

## BillingGroups Api

### Retrieve
```bash
curl https://api.lob.com/v1/billing_groups/bg_4bb02b527a72667d0 \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

fetchedBillingGroup, _, err := apiClient.BillingGroupsApi.Get(context,"bg_4bb02b527a72667d0").Execute()

if err != nil {
    return err
}
```




### Update
```bash
curl https://api.lob.com/v1/billing_groups/bg_4bb02b527a72667d0 \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=demo replacement" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var bgEditable = *lob.NewBillingGroupEditable("Test Billing group")
createdBillingGroup, _, _ := apiClient.BillingGroupsApi.Create(context).BillingGroupEditable(suite.bgEditable).Execute()

bgEditable.SetDescription("BG updated")
resp, _, err := suite.apiClient.BillingGroupsApi.Update(context, createdBillingGroup.Id).BillingGroupEditable(suite.bgEditable).Execute()

if err != nil {
    return err
}
```


### Create
```bash
curl https://api.lob.com/v1/billing_groups \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Usage group used for the Marketing Department's resource sends" \
  -d "name=Marketing Department" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)


var bgCreate = *lob.NewBillingGroupEditable()
bgCreate.SetDescription("Usage group used for the Marketing Department's resource sends")
bgCreate.SetName("Marketing Department")



createdbg, _, err := apiClient.BillingGroupsApi.Create(context).BillingGroupEditable(bgCreate).Execute()

if err != nil {
    return err
}
```


### Update
```bash
curl https://api.lob.com/v1/billing_groups/bg_759954f540a1bfdb5 \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=demo replacement" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var bgEditable = *lob.NewBillingGroupEditable("Test Billing group")
createdBillingGroup, _, _ := apiClient.BillingGroupsApi.Create(context).BillingGroupEditable(suite.bgEditable).Execute()

bgEditable.SetDescription("BG updated")
resp, _, err := suite.apiClient.BillingGroupsApi.Update(context, createdBillingGroup.Id).BillingGroupEditable(suite.bgEditable).Execute()

if err != nil {
    return err
}
```




### List
```bash
curl -X GET "https://api.lob.com/v1/billing_groups?limit=2" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)
BillingGroupList = apiClient.BillingGroupsApi.List(context).Execute()
if err != nil {
    return err
}
```
### Update
```bash
curl https://api.lob.com/v1/billing_groups/ \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=demo replacement" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var bgEditable = *lob.NewBillingGroupEditable("Test Billing group")
createdBillingGroup, _, _ := apiClient.BillingGroupsApi.Create(context).BillingGroupEditable(suite.bgEditable).Execute()

bgEditable.SetDescription("BG updated")
resp, _, err := suite.apiClient.BillingGroupsApi.Update(context, createdBillingGroup.Id).BillingGroupEditable(suite.bgEditable).Execute()

if err != nil {
    return err
}
```

## Cards Api

### Retrieve
```bash
curl https://api.lob.com/v1/cards/card_7a6d73c5c8457fc \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

fetchedCard, _, err := apiClient.CardsApi.Get(context,"card_7a6d73c5c8457fc").Execute()

if err != nil {
    return err
}
```




### Update
```bash
curl -X POST https://api.lob.com/v1/cards/card_7a6d73c5c8457fc \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Awesome card" \
  -d "auto_reorder=true"
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

cardEditable = *lob.NewCardEditable("https://s3-us-west-2.amazonaws.com/public.lob.com/assets/card_horizontal.pdf()")
createdCard, _, _ := apiClient.CardsApi.Create(context).CardEditable(suite.cardEditable).Execute()
var updatedCard = *lob.NewCardUpdatable()
updatedCard.SetDescription("Card was updated")

resp, _, err := suite.apiClient.CardsApi.Update(context, createdCard.Id).CardUpdatable(updatedCard).Execute()

if err != nil {
    return err
}
```



### Delete
```bash
curl -X DELETE "https://api.lob.com/v1/cards/card_6afffd19045076c" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

deletedCard, _, err := apiClient.CardsApi.Delete(context, "card_6afffd19045076c").Execute()

if err != nil {
    return err
}
```


### Update
```bash
curl -X POST https://api.lob.com/v1/cards/card_6afffd19045076c \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Awesome card" \
  -d "auto_reorder=true"
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

cardEditable = *lob.NewCardEditable("https://s3-us-west-2.amazonaws.com/public.lob.com/assets/card_horizontal.pdf()")
createdCard, _, _ := apiClient.CardsApi.Create(context).CardEditable(suite.cardEditable).Execute()
var updatedCard = *lob.NewCardUpdatable()
updatedCard.SetDescription("Card was updated")

resp, _, err := suite.apiClient.CardsApi.Update(context, createdCard.Id).CardUpdatable(updatedCard).Execute()

if err != nil {
    return err
}
```


### Create
```bash
curl https://api.lob.com/v1/cards \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "front=https://s3-us-west-2.amazonaws.com/public.lob.com/assets/card_horizontal.pdf" \
  -d "back=https://s3-us-west-2.amazonaws.com/public.lob.com/assets/card_horizontal.pdf" \
  -d "size=2.125x3.375" \
  -d "description=Test Card" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)


var cardCreate = *lob.NewCardEditable()
cardCreate.SetFront("https://s3-us-west-2.amazonaws.com/public.lob.com/assets/card_horizontal.pdf")
cardCreate.SetBack("https://s3-us-west-2.amazonaws.com/public.lob.com/assets/card_horizontal.pdf")
cardCreate.SetSize("2.125x3.375")
cardCreate.SetDescription("Test Card")



createdcard, _, err := apiClient.CardsApi.Create(context).CardEditable(cardCreate).Execute()

if err != nil {
    return err
}
```


### Update
```bash
curl -X POST https://api.lob.com/v1/cards/ \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Awesome card" \
  -d "auto_reorder=true"
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

cardEditable = *lob.NewCardEditable("https://s3-us-west-2.amazonaws.com/public.lob.com/assets/card_horizontal.pdf()")
createdCard, _, _ := apiClient.CardsApi.Create(context).CardEditable(suite.cardEditable).Execute()
var updatedCard = *lob.NewCardUpdatable()
updatedCard.SetDescription("Card was updated")

resp, _, err := suite.apiClient.CardsApi.Update(context, createdCard.Id).CardUpdatable(updatedCard).Execute()

if err != nil {
    return err
}
```




### List
```bash
curl -X GET "https://api.lob.com/v1/cards?limit=2" \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)
CardList = apiClient.CardsApi.List(context).Execute()
if err != nil {
    return err
}
```
### Update
```bash
curl -X POST https://api.lob.com/v1/cards/ \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Awesome card" \
  -d "auto_reorder=true"
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

cardEditable = *lob.NewCardEditable("https://s3-us-west-2.amazonaws.com/public.lob.com/assets/card_horizontal.pdf()")
createdCard, _, _ := apiClient.CardsApi.Create(context).CardEditable(suite.cardEditable).Execute()
var updatedCard = *lob.NewCardUpdatable()
updatedCard.SetDescription("Card was updated")

resp, _, err := suite.apiClient.CardsApi.Update(context, createdCard.Id).CardUpdatable(updatedCard).Execute()

if err != nil {
    return err
}
```





### Update
```bash
curl -X POST https://api.lob.com/v1/cards/card_6afffd19045076c \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Awesome card" \
  -d "auto_reorder=true"
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

cardEditable = *lob.NewCardEditable("https://s3-us-west-2.amazonaws.com/public.lob.com/assets/card_horizontal.pdf()")
createdCard, _, _ := apiClient.CardsApi.Create(context).CardEditable(suite.cardEditable).Execute()
var updatedCard = *lob.NewCardUpdatable()
updatedCard.SetDescription("Card was updated")

resp, _, err := suite.apiClient.CardsApi.Update(context, createdCard.Id).CardUpdatable(updatedCard).Execute()

if err != nil {
    return err
}
```

## CardOrders Api

### Retrieve
```bash
curl https://api.lob.com/v1/cards/card_6afffd19045076c/orders/ \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc:
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

fetchedCardOrderList, _, err := apiClient.CardOrdersApi.Get(context, "card_6afffd19045076c").Execute()

if err != nil {
    return err
}
```





### Create
```bash
curl https://api.lob.com/v1/cards/card_6afffd19045076c/orders \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "quantity=10000" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)


var cardOrderCreate = *lob.NewCardOrderEditable()
cardOrderCreate.SetQuantity("10000")



createdcardOrder, _, err := apiClient.CardOrdersApi.Create(context, "card_6afffd19045076c").CardOrderEditable(cardOrderCreate).Execute()

if err != nil {
    return err
}
```


## ZipLookups Api

### ZipLookup
```bash
curl https://api.lob.com/v1/us_zip_lookups \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "zip_code=94107"
```

```go
var context = context.Background()
context = context.WithValue(context, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var zipEditable = *lob.NewZipEditable()
zipEditable.SetZipCode("94107")

resp, _, err := apiClient.ZipLookupsApi.Lookup(context).ZipEditable(zipEditable).Execute()

if err != nil {
    return err
}
```

## Reverse Geocode Lookups Api

### Reverse Geocode Lookup
```bash
curl https://api.lob.com/v1/us_reverse_geocode_lookups \
  -u <YOUR_LIVE_API_KEY>: \
  -d "latitude=37.777456" \
  -d "longitude=-122.393039" \
```

```go
var context = context.Background()
context = context.WithValue(context, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var location = *lob.NewLocationWithDefaults()
location.SetLatitude(37.777456)
location.SetLongitude(122.393039)

resp, _, err := apiClient.ReverseGeocodeLookupsApi.Lookup(context).Location(location).Execute()

if err != nil {
    return err
}
```

## USAutoCompletions Api

### Autocomplete
```bash
curl https://api.lob.com/v1/us_autocompletions \
  -u <YOUR_LIVE_API_KEY>: \
  -d "address_prefix=185 B" \
  -d "city=San Francisco" \
  -d "state=CA" \
  -d "zip_code=94017" \
```

```go
var context = context.Background()
context = context.WithValue(context, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var usAutocomplationsWritable = *lob.NewUsAutocompletionsWritable("340 Wat")

resp, _, err := suite.apiClient.UsAutocompletionsApi.Autocomplete(suite.ctx).UsAutocompletionsWritable(suite.usAutocomplationsWritable).Execute()

if err != nil {
    return err
}
```

## IntlAutocompletion Api

### IntlAutocompletion
```bash
curl https://api.lob.com/v1/intl_autocompletions \
  -u <YOUR_LIVE_API_KEY>: \
  -d "address_prefix=340 Wat" \
  -d "city=Summerside" \
  -d "state=Prince Edward Island" \
  -d "zip_code=C1N 1C4" \
  -d "country=CA" \
```

```go
var context = context.Background()
context = context.WithValue(context, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

resp, _, err := suite.apiClient.IntlAutocompletionsApi.Autocomplete(context).IntlAutocompletionsWritable(suite.intlAutocomplationsWritable).Execute()

if err != nil {
    return err
}
```
