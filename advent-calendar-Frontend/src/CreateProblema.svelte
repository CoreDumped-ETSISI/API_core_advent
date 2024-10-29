<!-- src/CreateProblema.svelte -->
<script>
    import { onMount } from 'svelte';
    import { navigate } from 'svelte-routing';

    let year = '';
    let dia = '';
    let titulo = '';
    let enunciado = '';
    let solucion = '';
    let fechaDesbloqueo = '';
    let fechaBloqueo = '';
    let message = '';

    const createProblema = async () => {
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
                fecha_desbloqueo: fechaDesbloqueo,
                fecha_bloqueo: fechaBloqueo,
            }),
        });

        if (response.ok) {
            message = 'Problema creado exitosamente!';
            // Optionally redirect to another page
            navigate('/'); // Redirect to homepage or admin dashboard
        } else {
            const errorData = await response.json();
            message = `Error: ${errorData.error || 'Error desconocido'}`;
        }
    };
</script>

<main>
    <h1>Crear Problema</h1>
    
    <input type="number" bind:value={year} placeholder="Año" required />
    <input type="number" bind:value={dia} placeholder="Día" required />
    <input type="text" bind:value={titulo} placeholder="Título" required />
    <textarea bind:value={enunciado} placeholder="Enunciado" required></textarea>
    <textarea bind:value={solucion} placeholder="Solución" required></textarea>
    <input type="datetime-local" bind:value={fechaDesbloqueo} placeholder="Fecha Desbloqueo" required />
    <input type="datetime-local" bind:value={fechaBloqueo} placeholder="Fecha Bloqueo" required />
    <button on:click={createProblema}>Crear Problema</button>
    <p>{message}</p>
</main>

<style>
    /* Add your styles here */
</style>
