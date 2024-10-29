<script>
    import { Link } from 'svelte-routing';
</script>

<main>
    <h1>404 - Page Not Found</h1>
    <p>Sorry, the page you are looking for does not exist.</p>
    <Link to="/">Go to Home</Link>
</main>

<style>
    main {
        text-align: center;
        padding: 50px;
    }
    h1 {
        font-size: 2rem;
        color: #e74c3c; /* You can choose a color of your choice */
    }
    /* a {
        color: #3498db;
        text-decoration: none;
    }
    a:hover {
        text-decoration: underline;
    } */
</style>
