<!-- src/Register.svelte -->
<script>
    import { navigate } from 'svelte-routing';

    let correo = '';
    let usuario = '';
    let contrasena = '';
    let message = '';

    const register = async () => {
        const response = await fetch('http://localhost:8000/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ correo, usuario, contrasena }),
        });

        if (response.ok) {
            message = 'Registro exitoso!';
            // Optionally redirect to login or another page
            navigate('/login'); // Redirect to login page
        } else {
            const errorData = await response.json();
            message = `Error: ${errorData.error || 'Error desconocido'}`;
        }
    };
</script>

<main>
    <h1>Registro de Usuario</h1>
    <p>{message}</p>
    <input type="email" bind:value={correo} placeholder="Correo" required />
    <input type="text" bind:value={usuario} placeholder="Usuario" required />
    <input type="password" bind:value={contrasena} placeholder="Contraseña" required />
    <button on:click={register}>Registrar</button>
</main>

<style>
    /* Add your styles here */
</style>
