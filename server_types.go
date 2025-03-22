<?php
  include("coneccion.php");
?>


<form method="post" autocomplete="off">

<form method="post">


<label for="numero1">Numero 1</label>
<input type="text" id="numero1" onchange="sumar();"><br>

<label for="numero2">Numero 2</label>
<input type="text" id="numero2" onchange="sumar();"><br>

<label for="resultado">Resultado</label>
<input type="text" id="resultado">
  




<input type="submit" value=suma"Enviar">
            
 <input type="reset" value="cancelar"> 
          
 
 </form>



<script>
    // si la respuesta que se espera es sumar
    function sumar(){
        var numero1 = document.getElementById('numero1').value;
        var numero2 = document.getElementById('numero2').value;

        if(numero1!=='' && numero2!==''){
            var suma = parseInt(numero1)+parseInt(numero2);
            document.getElementById('resultado').value = suma;

     




        } 
    }
   

</script>

<?php

 if(isset($_POST['Enviar'])){
      
    $numero1 = $_POST['numero1'];
    $numero2 = $_POST['numero2'];
    $resultado = $_POST['resultado'];
    
      $fecha = date("d/m/y");

      $insertar = "INSERT INTO datosuma Values ('$numero1','$numero2','$resultado','$fecha','')";
      
      $conex = mysqli_query($coneccion,$insertar);
  }

?>
