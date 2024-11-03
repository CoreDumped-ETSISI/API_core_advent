En la búsqueda del componente pérdido

Después del asalto a la rotonda, los CoreElfos tienen que construir regalos muy rápido, en las prisas han perdido los componentes de los regalos. Cada regalo tiene un componente que se ha perdido, pero no saben cuál es.

Pero los CoreElfos cometen errores muy matemáticos, sabemos que la regla con la que han perdido los componentes es la siguiente:
- Cada componente es una letra, y tiene un valor (a=1, b=2, c=3, ..., z=26, A = 27, B = 28, ..., Z = 52).
- Cada regalo es una mezcla de números y letras.
- El primer número de cada regalo es la posición relativa del componente perdido (dado "a2bcda", el componente perdido será 'c' porque es el segundo componente a partir del cáracter '2').
- En los casos de juguetes muy complejos, el dígito puede apuntar a otro dígito, que a su vez puede apuntar a más dígitos o a la letra que buscamos (dado "a2v3e2gfh", el componente perdido será 'g' porque es el tercer componente a partir del cáracter '3', que es el segundo a partir del cáracter '2').

Dada esta lista de regalos, encuentra la suma de los valores de todos los componentes perdidos.