'use strict';

angular.module('angappApp')
  .controller('PersonsCtrl', function ($scope, personsService) {
    var getPersonsSuccess = function(response){
    	$scope.persons = response.Value;
    };

    var getPersonsError = function(response){

    };

    $scope.refresh = function(){
    	personsService.getPersons(getPersonsSuccess, getPersonsError);
    };

    $scope.refresh();


  });
