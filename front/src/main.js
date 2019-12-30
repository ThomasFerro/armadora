import App from './App.svelte';
import { createGame } from './gameFactory';

const app = new App({
	target: document.body,
	props: {
		game: createGame()
	}
});

export default app;
