<!-- src/Login.svelte -->
<script>
    import { onMount } from 'svelte';
    import { navigate } from 'svelte-routing';

    let valor = '';
    let contrasena = '';
    let message = '';

    const login = async () => {
        const response = await fetch('http://localhost:8000/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ valor, contrasena }),
        });

        if (response.ok) {
            const data = await response.json();
            localStorage.setItem('token', data.token); // Store the token for future requests
            message = 'Inicio de sesión exitoso!';
            // Optionally redirect to a different page
            navigate('/'); // Redirect to homepage or admin page
        } else {
            const errorData = await response.json();
            message = `Error: ${errorData.error || 'Error desconocido'}`;
        }
    };
</script>

<main>
    <h1>Inicio de Sesión</h1>
    <input type="text" bind:value={valor} placeholder="Correo o Usuario" required />
    <input type="password" bind:value={contrasena} placeholder="Contraseña" required />
    <button on:click={login}>Iniciar Sesión</button>
    <p>{message}</p>
</main>

<style>
    /* Add your styles here */
</style>
