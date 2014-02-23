'use strict';

angular.module('angappApp')
  .controller('AboutCtrl', function ($scope, aboutService) {

  	var getInfoSuccess = function(data){
  		if (data && data.aboutus){
  			$scope.aboutInfo = data.aboutus;
  		}
  	};

  	var getInfoError = function(){
  		console.log("crud we errored out man.");
  	};
  	
  	$scope.refresh = function(){
	    aboutService.getAboutInfo(getInfoSuccess, getInfoError);
	  };

	  $scope.refresh();
  });
