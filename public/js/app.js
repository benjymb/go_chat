var app = angular.module('universidadApp', []);
app.controller('profesorCtrl', function($scope){
	$scope.profesor = profesorData;
	$scope.editando = {};
	$scope.mostrarForm = false;
	$scope.EditarProfesor = function(){
		angular.copy($scope.profesor, $scope.editando);
		$scope.mostrarForm = true;
	}
	$scope.GuardarCambios = function(){
		$scope.profesor = $scope.editando;
		$scope.editando = {};
		$scope.mostrarForm = false;
	}
	$scope.CancelarCambios = function(){
		$scope.editando = {};
		$scope.mostrarForm = false;
	}
})

var profesorData = {
        nombre: "Juan Carlos Pineda",
        bio: "Saludos querido estudiante, mi nombre es Juan Carlos y soy su profesor.",
        edad: 47,
        foto: "img/juancarlos.jpg"
}
	
