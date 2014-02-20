angular.module('angappApp')
	.factory('aboutService', function(){
		var factory = {};

		factory.getAboutInfo = function(){
			return "about us information; this might instead come from an API call.";
		};

		return factory;
	});