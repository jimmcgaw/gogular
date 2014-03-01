'use strict';

angular.module('angappApp')
  .controller('EditPersonsCtrl', function ($scope, $routeParams, personsService) {
    var updatePersonSuccess = function(response){
    	// what should I do here?
    };

    var updatePersonError = function(response){

    };

    $scope.updatePerson = function(person){
    	var params = {
    		'first_name' : person.FirstName,
    		'last_name' : person.LastName,
        'personId' : $routeParams.id
    	};
    	personsService.updatePerson(params, updatePersonSuccess, updatePersonError);
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
