'use strict';

angular.module('angappApp')
  .controller('AboutCtrl', function ($scope, aboutService) {
    $scope.aboutInfo = aboutService.getAboutInfo();
  });
