# Getting Started

This guide illustrates how to use the SDK to implement the following method patterns:

- CREATE
- LIST
- GET
- DELETE
- VERIFY (BANK ACCOUNTS)
- UPDATE

## INSTALL

```
$ go install github.com/lob/lob-go@latest
```

## IMPORT AND INITIALIZE

```go
import (
  "context"
  "os"
  lob "github.com/lob/lob-go"
)
ctx := context.Background()
ctx = context.WithValue(ctx, lob.ContextBasicAuth, lob.BasicAuth{UserName: os.Getenv("LOB_API_TEST_KEY")})

config := *lob.NewConfiguration()
```
You then instantiate the specific resource API that you need access to as follows:

```go
apiClient := *lob.NewAPIClient(&config)
```

## METHODS

This abstracts the request/response interaction out of the calling code such that you create the resource and pass it to the corresponding API using error handling.

### CREATE METHOD

Here is a sample of the CREATE method

```go
address := *lob.AddressEditable
address.SetName("Harry Zhang")
address.SetAddressLine1("2261 Market Street")
address.SetAddressCity("San Francisco")
address.SetAddressState("CA")
address.SetAddressZip("94114")

resp, _, err := apiClient.AddressesApi.Create(ctx).AddressEditable(address).Execute()

if err != nil {
    log.Fatal(err)
}
```

### LIST METHOD

Here is a sample of the LIST method:

```go
resp, _, err := apiClient.AddressesApi.List(ctx).Limit(3).Execute()

if err != nil {
    log.Fatal(err)
}
```

### GET BY ID METHOD

Here is a sample of the GET method:

```go
resp, _, err := apiClient.AddressesApi.Get(ctx, "adr_xxx").Execute()

if err != nil {
    log.Fatal(err)
}
```

### DELETE METHOD

Here is a sample of the DELETE method:

```go
resp, _, err := apiClient.AddressesApi.Delete(ctx, "adr_xxx").Execute()

if err != nil {
    log.Fatal(err)
}
```

### BANK ACCOUNT VERIFY

Here is a sample of the BANK ACCOUNT VERIFY method:

```go
verifyAmounts := []int32{11, 35}
bankAccountVerify := *lob.NewBankAccountVerify(verifyAmounts)

resp, _, err := apiClient.BankAccountsApi.Verify(ctx, "bank_xxx").BankAccountVerify(bankAccountVerify).Execute()

if err != nil {
    log.Fatal(err)
}
```

### UPDATE METHOD

Here is a sample of the UPDATE method:

```go
templateUpdate := *lob.TemplateUpdate
templateUpdate.SetPublishedVersion("vrsn_xxx")
templateUpdate.SetDescription("Updated Template")

resp, _, err := apiClient.TemplatesApi.Update(ctx, "tmpl_xxx").TemplateUpdate(templateUpdate).Execute()

if err != nil {
    log.Fatal(err)
}
```
