import { Route, BrowserRouter as Router, Routes } from "react-router-dom";
import Home from "./pages/Home";
import TicketPortal from "./pages/Ticket";
import Issuer from "./pages/Issuer";

function App() {
  return (
    <Router>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/issuer" element={<Issuer />} />
            <Route path="/ticket" element={<TicketPortal />} />
          </Routes>
        </Router>
  );
}

export default App;
