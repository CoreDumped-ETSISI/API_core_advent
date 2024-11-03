<script>
    import { goto } from '$app/navigation'; // Import goto from SvelteKit

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
            // Redirect to login page
            goto('/login'); // Use goto for navigation
        } else {
            const errorData = await response.json();
            message = `Error: ${errorData.error || 'Error desconocido'}`;
        }
    };
</script>

<main>
    <h1>Registro de Usuario</h1>
    <p>{message}</p>
    <form on:submit|preventDefault={register}> <!-- Prevent default form submission -->
        <input type="email" bind:value={correo} placeholder="Correo" required />
        <input type="text" bind:value={usuario} placeholder="Usuario" required />
        <input type="password" bind:value={contrasena} placeholder="Contraseña" required />
        <button type="submit">Registrar</button> <!-- Change to type="submit" -->
    </form>
</main>

<style>
    main {
        text-align: center;
        padding: 20px;
    }
    h1 {
        font-size: 2rem;
        color: #3498db; /* Choose your color */
    }
    input {
        margin: 10px 0;
        padding: 10px;
        width: 100%;
        max-width: 300px; /* Limit width for better usability */
        border: 1px solid #ccc;
        border-radius: 4px;
    }
    button {
        padding: 10px 15px;
        background-color: #3498db;
        color: white;
        border: none;
        border-radius: 5px;
        cursor: pointer;
    }
    button:hover {
        background-color: #2980b9;
    }
</style>
