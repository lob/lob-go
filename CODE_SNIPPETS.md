# Code Snippets

## Address Api





### Create
```bash
curl https://api.lob.com/v1/addresses \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Harry - Office" \
  -d "name=Harry Zhang" \
  -d "company=Lob" \
  -d "email=harry@lob.com" \
  -d "phone=5555555555" \
  -d "address_line1=210 King St" \
  -d "address_line2=# 6100" \
  -d "address_city=San Francisco" \
  -d "address_state=CA" \
  -d "address_zip=94107" \
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
addressCreate.SetAddressLine1("210 King St")
addressCreate.SetAddressLine2("# 6100")
addressCreate.SetAddressCity("San Francisco")
addressCreate.SetAddressState("CA")
addressCreate.SetAddressZip("94107")
addressCreate.SetAddressCountry("US")



createdaddress, _, err := apiClient.AddressesApi.Create(context).AddressEditable(addressCreate).Execute()

if err != nil {
    return err
}
```























## Postcards Api





### Create
```bash
curl https://api.lob.com/v1/postcards \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Demo Postcard job" \
  -d "from=adr_210a8d4b0b76d77b" \
  --data-urlencode "front=<html style='padding: 1in; font-size: 50;'>Front HTML for {{name}}</html>" \
  --data-urlencode "back=<html style='padding: 1in; font-size: 20;'>Back HTML for {{name}}</html>" \
  -d "to[name]=Harry Zhang" \
  -d "to[address_line1]=210 King St" \
  -d "to[address_line2]=# 6100" \
  -d "to[address_city]=San Francisco" \
  -d "to[address_state]=CA" \
  -d "to[address_zip]=94107" \
  -d "merge_variables[name]=Harry" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var to = *lob.NewAddressEditable()
to.SetAddressLine1("210 King St")
to.SetAddressLine2("# 6100")
to.SetAddressCity("San Francisco")
to.SetAddressState("CA")
to.SetAddressZip("94107")
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



## SelfMailers Api





### Create
```bash
curl https://api.lob.com/v1/self_mailers \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Demo Self Mailer job" \
  -d "from=adr_210a8d4b0b76d77b" \
  --data-urlencode "inside=<html style='padding: 1in; font-size: 50;'>Inside HTML for {{name}}</html>" \
  --data-urlencode "outside=<html style='padding: 1in; font-size: 20;'>Outside HTML for {{name}}</html>" \
  -d "to[name]=Harry Zhang" \
  -d "to[address_line1]=210 King St" \
  -d "to[address_line2]=# 6100" \
  -d "to[address_city]=San Francisco" \
  -d "to[address_state]=CA" \
  -d "to[address_zip]=94107" \
  -d "merge_variables[name]=Harry" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var to = *lob.NewAddressEditable()
to.SetAddressLine1("210 King St")
to.SetAddressLine2("# 6100")
to.SetAddressCity("San Francisco")
to.SetAddressState("CA")
to.SetAddressZip("94107")
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



## Letters Api





### Create
```bash
curl https://api.lob.com/v1/letters \
  -u test_0dc8d51e0acffcb1880e0f19c79b2f5b0cc: \
  -d "description=Demo Letter" \
  -d "from=adr_210a8d4b0b76d77b" \
  --data-urlencode "file=<html style='padding-top: 3in; margin: .5in;'>HTML Letter for {{name}}</html>" \
  -d "color=true" \
  -d "to[name]=Harry Zhang" \
  -d "to[address_line1]=210 King St" \
  -d "to[address_line2]=# 6100" \
  -d "to[address_city]=San Francisco" \
  -d "to[address_state]=CA" \
  -d "to[address_zip]=94107" \
  -d "merge_variables[name]=Harry" \
  -d "cards[]=card_c51ae96f5cebf3e"
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var to = *lob.NewAddressEditable()
to.SetAddressLine1("210 King St")
to.SetAddressLine2("# 6100")
to.SetAddressCity("San Francisco")
to.SetAddressState("CA")
to.SetAddressZip("94107")
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



## Checks Api





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
  -d "to[address_line1]=210 King St" \
  -d "to[address_line2]=# 6100" \
  -d "to[address_city]=San Francisco" \
  -d "to[address_state]=CA" \
  -d "to[address_zip]=94107" \
  -d "merge_variables[name]=Harry" \
```

```go
var context = context.Background()
context = context.WithValue(suite.ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("<YOUR_API_KEY>")})

var apiClient = *lob.NewAPIClient(configuration)

var to = *lob.NewAddressEditable()
to.SetAddressLine1("210 King St")
to.SetAddressLine2("# 6100")
to.SetAddressCity("San Francisco")
to.SetAddressState("CA")
to.SetAddressZip("94107")
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



## BankAccounts Api









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
templateCreate.Set("Test Template")
templateCreate.Set("<html>HTML for {{name}}</html>")



createdtemplate, _, err := apiClient.TemplatesApi.Create(context).TemplateWritable(templateCreate).Execute()

if err != nil {
    return err
}
```

## TemplateVersions Api









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
templateVersionCreate.Set("Second Version")
templateVersionCreate.Set("<html>Second HTML for {{name}}</html>")



createdtemplateVersion, _, err := apiClient.TemplateVersionsApi.Create(context, "tmpl_4aa14648113e45b").TemplateVersionWritable(templateVersionCreate).Execute()

if err != nil {
    return err
}
```

