# GQL - Update Ticket
[Back](../README.md#updating-ticket-status-flow)

**Query**: UpdateTicket

**Description**: Update specific ticket status based on id

## Parameter

| Field  | Type                  | Description      |
| ------ | --------------------- | ---------------- |
| id     | int32 **(required)**  | Target ticket id |
| status | string **(required)** | Status change    |

## Example Request

```
mutation UpdateTicket {
    updateTicket(
        input: { 
                id: 1, 
                status: "Completed" 
        }
    ) 
    {
        id
        success
    }
}
```