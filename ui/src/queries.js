import gql from 'graphql-tag';

export const CHATS = gql`
query chats {
	chats {
		id
		title
		last_changed
	}
}
`;


export const CHAT = gql`
query chat($id: ID!) {
	chat(id: $id) {
		id
		model
		messages {
			role
			content
		}
	}
}
`;

export const SEND = gql`
subscription send($id: ID!, $model: String!, $content: String!) {
	send(id: $id, model: $model content: $content) {
		message {
			content
		}
	}
}`;
