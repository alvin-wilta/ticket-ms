import { gql } from 'apollo-server';

const typeDefs = gql`
type Ticket {
  id: Int
  title: String!
  description: String
  status: String
  name: String!
}

input TicketFilter {
  id: Int
  status: String
}

input NewTicket {
  title: String!
  description: String!
  name: String!
}

type CreateTicketResponse {
  success: Boolean!
}

input UpdateTicket {
  id: Int!
  status: String!
}

type UpdateTicketResponse {
  id: Int!
  success: Boolean!
}

input DeleteTicket {
  id: Int!
}

type DeleteTicketResponse {
  id: Int!
  success: Boolean!
}

type Query {
  healthCheck: String!
  tickets(input: TicketFilter): [Ticket!]!
}

type Mutation {
  createTicket(input: NewTicket!): CreateTicketResponse!
  updateTicket(input: UpdateTicket!): UpdateTicketResponse!
  deleteTicket(input: DeleteTicket!): DeleteTicketResponse!
}

`