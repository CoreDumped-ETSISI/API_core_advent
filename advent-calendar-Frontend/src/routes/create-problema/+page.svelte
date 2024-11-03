<script>
    import { goto } from '$app/navigation'; // Import goto from SvelteKit

    let year = '';
    let dia = '';
    let titulo = '';
    let enunciado = '';
    let solucion = '';
    let fechaDesbloqueo = '';
    let fechaBloqueo = '';
    let message = '';


    
    const createProblema = async () => {
        const formatDate = (dateString) => {
            const dateObj = new Date(dateString);
            return dateObj.toISOString();
        };
        const response = await fetch('http://localhost:8000/admin/problemas', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('token')}`, // Include the token if needed
            },
            body: JSON.stringify({
                year: parseInt(year),
                dia: parseInt(dia),
                titulo,
                enunciado,
                solucion,
                fecha_desbloqueo: formatDate(fechaDesbloqueo),
                fecha_bloqueo: formatDate(fechaBloqueo),
            }),
        });

        if (response.ok) {
            message = 'Problema creado exitosamente!';
            // Redirect to another page (e.g., homepage or admin dashboard)
            goto('/'); // Use goto for navigation
        } else {
            const errorData = await response.json();
            message = `Error: ${errorData.error || 'Error desconocido'}`;
        }
    };
</script>

<main>
    <h1>Crear Problema</h1>

    <form on:submit|preventDefault={createProblema}> <!-- Prevent default form submission -->
        <input type="number" bind:value={year} placeholder="Año" required />
        <input type="number" bind:value={dia} placeholder="Día" required />
        <input type="text" bind:value={titulo} placeholder="Título" required />
        <textarea bind:value={enunciado} placeholder="Enunciado" required></textarea>
        <textarea bind:value={solucion} placeholder="Solución" required></textarea>
        <input type="datetime-local" bind:value={fechaDesbloqueo} placeholder="Fecha Desbloqueo" required />
        <input type="datetime-local" bind:value={fechaBloqueo} placeholder="Fecha Bloqueo" required />
        <button type="submit">Crear Problema</button> <!-- Use type="submit" -->
    </form>

    <p>{message}</p>
</main>

<style>
    main {
        text-align: center;
        padding: 20px;
    }
    h1 {
        font-size: 2rem;
        color: #3498db; /* Customize your color */
    }
    input, textarea {
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
