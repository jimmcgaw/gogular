angular.module('angappApp')
	.factory('personsService', function($resource){

		var personsResource = $resource('/api/persons', {}, {
			'query' : {
				'method' : 'GET',
				'isArray' : false
			}
		});

		var personResource = $resource('/api/persons/:personId', { personId : '@id'}, {
			'get' : { 
				'method' : 'GET',
				'isArray' : false
			}
		});

		var factory = {};

		factory.getPersons = function(success, error){
			return personsResource.query(success, error)
		};

		factory.getPerson = function(params, success, error) {
			return personResource.get(params, success, error);
		}

		return factory;
	});