import { Button, Card, Container } from "react-bootstrap";

const Home = () => {

  return (
    <Container className="d-flex justify-content-center align-items-center text-center">
      <Card
        style={{ width: "25vw" }}
        className="mt-5 text-white"
        bg="dark"
        border="secondary"
      >
        <Card.Body>
          <Card.Header>
            <h2 className="fw-bold">Toped Support</h2>
          </Card.Header>
          <Card.Subtitle className="mt-2 mb-2">
            A simple ticket support app.
          </Card.Subtitle>
          <Container className="d-flex justify-content-center align-items-center">
            <Button href="/issuer" variant="primary" size="lg" className="m-3">
              New Issue
            </Button>
            <br />
            <Button href="/ticket" variant="secondary" size="lg">
              Ticket List
            </Button>
          </Container>
        </Card.Body>
      </Card>
    </Container>
  );
};

export default Home;
