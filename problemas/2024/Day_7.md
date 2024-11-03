Robos en caliente

El servidor central de CoreClaus ha sido hackeado. Han parado el proceso que calculaba el viaje de CoreClaus por el mundo. Ese proceso estaba repartido entre varias salas de servidores. Han atacado un solo servidor y han tenido que hacerlo en persona, al acabar el proceso (entrenamientos de modelos de LLM muy complejos) se ha parado el servidor, por lo que la temperatura de su sala debe haber caido durante unos momentos. No sabemos que sala ha sido atacada ni cuando, pero tenemos un sistema de seguridad que nos proporciona la siguiente información en tablas SQL:

- Sala(id) => Un registro de las salas que contienen el servidor.
- Temperatura(id, temperatura, tiempo, id_sala) => Un registro de la temperatura de las salas en cada minuto.
- Cámara(id, id_sala) => Un registro de las cámaras de seguridad en cada sala.
- PersonasReconocidad(id, id_cámara, nombre, tiempo) => Un registro de las personas reconocidas por las cámaras en cada minuto.

Queremos saber el nombre del sospechoso. 