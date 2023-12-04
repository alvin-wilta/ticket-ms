import { Button, Card, Container } from "react-bootstrap";
import { Ticket } from "../types/ticket";


const TicketCard = (ticket: Ticket) => {
  // const navigate = useNavigate();

  // const handleDetailClick = () => {
  //   console.log("navigate");
  //   navigate(`/jobs/${job.id}`);
  // };
  return (
    <Card style={{ width: "22rem" }} className="m-3 d-inline-block">
      <Card.Body>
        <Card.Title className="fw-bold">{ticket.title}</Card.Title>
        <Card.Subtitle>
          Issuer: {ticket.name} - {ticket.status}
        </Card.Subtitle>
        <Card.Text className="text-secondary mt-1">
          {ticket.description}
        </Card.Text>
      </Card.Body>
    </Card>
  );
};

export default TicketCard;
