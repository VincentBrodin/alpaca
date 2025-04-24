import { ApolloClient, HttpLink, split, InMemoryCache } from '@apollo/client/core'
import { GraphQLWsLink } from '@apollo/client/link/subscriptions'
import { createClient } from 'graphql-ws'
import { getMainDefinition } from '@apollo/client/utilities';


const httpLink = new HttpLink({
	uri: 'http://localhost:1345/query'
})

const wsLink = new GraphQLWsLink(
	createClient({
		url: 'ws://localhost:1345/query',
	})
)

const link = split(
	({ query }) => {
		const definition = getMainDefinition(query)
		return (
			definition.kind === 'OperationDefinition'
			&& definition.operation === 'subscription'
		)
	},
	wsLink,
	httpLink
)

const cache = new InMemoryCache()

const apolloClient = new ApolloClient({
	link: link,
	cache,
})

export default apolloClient;

