var warningTrackApp = angular.module('warningTrackApp', []);

angular.module('warningTrackApp')
  .filter('svgIconBaseHref', function ($sce) {
    return function(basesId) {
      var brId = 0;
      if (isNaN(basesId)) {
        brId = basesId;
      }
      return $sce.trustAsResourceUrl('baseball-bases.svg#br' + brId);
    };
  });

angular.module('warningTrackApp')
  .filter('svgIconStrikesOutsHref', function ($sce) {
    return function(soId) {
      return $sce.trustAsResourceUrl('strikes-outs.svg#so' + soId);
    };
  });

angular.module('warningTrackApp')
  .filter('svgIconBallsHref', function ($sce) {
    return function(ballsId) {
      return $sce.trustAsResourceUrl('balls.svg#b' + ballsId);
    };
  });

warningTrackApp.controller('WarningTrackCtrl', ['$scope', '$http', '$filter',
  function ($scope, $http, $filter) {
    $http.get('/games').success(function(data) {
      $scope.games = data;
    });
    $scope.orderProp = ['status.status','-leverage_index'];
  }]);
