'use strict';

angular.module('angappApp')
  .controller('PersonsShowCtrl', function ($scope, $routeParams, personsService) {
    var getPersonSuccess = function(response){
      console.log(response.Value);
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
