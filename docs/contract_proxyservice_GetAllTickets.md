# GQL - Get All Tickets
[Back](../README.md#getting-ticket-flow)

**Query**: Tickets

**Description**: Get all tickets with certain parameter (status, name)

## Parameter

| Field  | Type   | Description                   |
| ------ | ------ | ----------------------------- |
| id     | int32  | Ticket id                     |
| status | string | Ticket status (New/Completed) |


## Example Request

```
query Tickets {
    tickets(input: { id: 1, status: "New" }) {
        id
        title
        description
        status
        name
    }
}
```