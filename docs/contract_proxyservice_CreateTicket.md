# GQL - Create Ticket
[Back](../README.md#creating-new-ticket-flow)

**Mutation**: CreateTicket

**Description**: Create new support ticket

## Parameter

| Field       | Type                  | Description           |
| ----------- | --------------------- | --------------------- |
| Title       | string **(required)** | ticket title          |
| Description | string **(required)** | ticket description    |
| Name        | string **(required)** | ticket requester name |

## Example Request

```
mutation CreateTicket {
    createTicket(
        input: { 
            title: "title", 
            description: "desc", 
            name: "issuer" 
        }) 
        {
            success
    }
}
```
