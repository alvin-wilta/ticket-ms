# GQL - Delete Ticket
[Back](../README.md#deleting-ticket-flow)

**Mutation**: DeleteTicket

**Description**: Delete a specific support ticket

## Parameter

| Field | Type  | Description |
| ----- | ----- | ----------- |
| id    | int32 | Ticket id   |

## Example Request

```
mutation DeleteTicket {
    deleteTicket(input: { id: 1 }) 
    {
        id
        success
    }
}
