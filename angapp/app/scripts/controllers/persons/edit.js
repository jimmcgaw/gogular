'use strict';

angular.module('angappApp')
  .controller('EditPersonsCtrl', function ($scope, $routeParams, personsService) {
    var savePersonSuccess = function(response){
    	// what should I do here?
    };

    var savePersonError = function(response){

    };

    $scope.savePerson = function(person){
    	var params = {
    		'first_name' : person.FirstName,
    		'last_name' : person.LastName
    	};
    	personsService.savePerson(params, savePersonSuccess, savePersonError);
    };

    var getPersonSuccess = function(response){
        $scope.person = response.Value;
    };

    var getPersonError = function(response){

    };

    $scope.refresh = function(){
      var params = {
        personId : $routeParams.id
      };
      personsService.getPerson(params, getPersonSuccess, getPersonError);
    };

    $scope.refresh();

  });
