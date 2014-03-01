angular.module('angappApp')
	.factory('personsService', function($resource){

		var personsResource = $resource('/api/persons', {}, {
			'query' : {
				'method' : 'GET',
				'isArray' : false
			},
			'save' : {
				'method' : 'POST',
				'isArray' : false,
				'params' : {
					'first_name' : '@first_name',
					'last_name' : '@last_name'
				}
			}
		});

		var personResource = $resource('/api/persons/:personId', { personId : '@id'}, {
			'get' : { 
				'method' : 'GET',
				'isArray' : false
			},
			'update' : {
				'method' : 'PUT',
				'isArray' : false,
				'params' : {
					'FirstName' : '@first_name', 
					'LastName' : '@last_name'
				}
			} 
		});

		var factory = {};

		factory.getPersons = function(success, error){
			return personsResource.query(success, error);
		};

		factory.getPerson = function(params, success, error) {
			return personResource.get(params, success, error);
		};

		factory.savePerson = function(params, success, error) {
			personsResource.save(params, success, error);
		};

		factory.updatePerson = function(params, success, error) {
			personResource.update(params, success, error);
		}

		return factory;
	});