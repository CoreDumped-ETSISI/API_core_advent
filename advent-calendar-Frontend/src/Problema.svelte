<script>
    import { onMount } from 'svelte';
    import { navigate } from 'svelte-routing';

    // export let year;
    // export let day;

    let problema = null;
    let error = '';
    let isAuthenticated = false;

    const checkAuthentication = () => {
        const token = localStorage.getItem('authToken');
        isAuthenticated = Boolean(token);
        return isAuthenticated;
    };

    const fetchProblema = async () => {
        if (!checkAuthentication()) {
            error = 'Debe estar autenticado para ver este contenido.';
            return;
        }

        try {
            const token = localStorage.getItem('authToken');
            const response = await fetch(`http://localhost:8000/${year}/${day}`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });
            if (!response.ok) {
                throw new Error('No se pudo cargar el problema');
            }

            problema = await response.json();
        } catch (err) {
            error = err.message || 'Ocurrió un error desconocido';
        }
    };

    // onMount(fetchProblema);
</script>

<main>
    <h1>Detalles del Problema</h1>

    <!-- {#if error}
        <p style="color: red;">{error}</p>
        <button on:click={() => navigate('/login')}>Iniciar sesión</button> -->
    <!-- {:else if problema} -->
        <!-- <h2>{problema.titulo}</h2> -->
        <!-- <p><strong>Año:</strong> {problema.year}</p>
        <p><strong>Día:</strong> {problema.dia}</p> -->
        <!-- <p><strong>Enunciado:</strong> {problema.enunciado}</p> -->
        <!-- <p><strong>Solución:</strong> {problema.solucion}</p>
        <p><strong>Fecha de Desbloqueo:</strong> {new Date(problema.fecha_desbloqueo).toLocaleString()}</p>
        <p><strong>Fecha de Bloqueo:</strong> {new Date(problema.fecha_bloqueo).toLocaleString()}</p> -->
    <!-- {:else}
        <p>Cargando los detalles del problema...</p>
    {/if} -->
</main>

<style>
    main {
        padding: 20px;
    }
    h1 {
        color: #333;
    }
</style>
