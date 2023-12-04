import { Card, Container } from "react-bootstrap";
import TicketCard from "../components/TicketCard";
import { Ticket } from "../types/ticket";

const ticketList: Ticket[] = [
  {
    title: "test",
    id: 1,
    description: "hello",
    status: "New",
    name: "Al"
  }
]

const TicketPortal = () => {
    return (
        <Container style={{ height: "50vh" }}>
          {ticketList.map((ticket) => (
            <TicketCard
              id={ticket.id}
              title={ticket.title}
              description={ticket.description}
              status={ticket.status}
              name={ticket.name}
            />
          ))}
        </Container>
      );
}

export default TicketPortal