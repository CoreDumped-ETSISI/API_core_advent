Navidades binarias

El compilador de los CoreDesarrolladoresDeNavidad ya no funciona (un becario hizo un push a producción sin pasar los tests). No da tiempo a hacer un revert (nadie sabe hacerlo porque no fueron al taller de git) y no hay copias de seguridad (el becario también se encargaba de eso); así que toca reescribir el compilador. Lo bueno es que tenemos el manual del MIPS128 (mucho mejor que el 64) y sabemos que el compilador original era muy sencillo. 

Sabemos que el compilador original solo soportaba 3 instrucciones: `ADD`, `SUB`, `MOV`, `SET` y `NOP`. 
- `ADD A B` suma dos registros y guarda el resultado en el primer registro.
- `SUB A B` resta dos registros y guarda el resultado en el primer registro.
- `MOV A B` copia el valor del segundo registro al primero.
- `SET A num` guarda el valor `num` en el registro A.
- `NOP A B` no hace nada.

Se codifican de la siguiente manera:
Los primeros 8 bits son la instrucción y los siguientes 16 bits son los operandos.
- `ADD` se codifica como un número impar que tiene más 1s que 0s.
- `SUB` se codifica como un número par que tiene más 0s que 1s.
- `MOV` se codifica como un número par que tiene más 1s que de 0s.
- `SET` se codifica como un número que tiene el mismo número de 0s que de 1s.
- `NOP` se codifica como un número impar que tiene más 0s que 1s.
Cada operando ocupa 8 bits.
- Un registro se codifica como una letra mayuscula en ASCII.
- Un número `num` se codifica como su binario sencillo (sin signo).
```
ADD A B -> 10111001 01000001 01000010
```
(Sabemos que el compilador original no tiene sentido, pero es lo que hay cuando el código es anterior a Alan Turing).

Teniendo el siguiente codigo compilado: 

Devuelve la palabra resultante de ejecutar el código y ordenar las letras de los registros de menor a mayor.

Por ejemplo:
```
O = 8
C = 5
E = 16
R = 12

Daría la palabra CORE.
```