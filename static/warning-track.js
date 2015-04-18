var warningTrackApp = angular.module('warningTrackApp', []);

warningTrackApp.controller('WarningTrackCtrl', ['$scope', '$http',
  function ($scope, $http) {
    $http.get('/games.json').success(function(data) {
      $scope.games = data;
    });

    $scope.orderProp = ['status','-leverage_index'];
  }]);
