<script>
	let firstName = "sdfg";
	let lastName = 'sdfasdf'
	let color = 'blue'

	$: name = firstName + ' ' + lastName
	const toggle = () => {color = color === 'blue' ? 'red' : 'blue'}
</script>

<main>
	<h1 style="color: {color}">Hello {name}!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
	<button on:click={toggle}>Purple</button>
</main>
