angular.module('angappApp')
	.factory('aboutService', function($resource){

		var aboutUsResource = $resource('/api/about', {}, {
			'query' : {
				'method' : 'GET',
				'isArray' : false
			}
		})

		var factory = {};

		factory.getAboutInfo = function(success, error){
			return aboutUsResource.query(success, error)
		};

		return factory;
	});