var warningTrackApp = angular.module('warningTrackApp', []);

warningTrackApp
  .filter('svgIconBaseHref', function($sce) {
    return function(basesId) {
      return $sce.trustAsResourceUrl('baseball-bases.svg#br' + basesId);
    };
  });

warningTrackApp
  .filter('displayGameStatus', function() {
    return function(game) {
      var displayString = game.status.status;
      if (game.status.status == "In Progress") {
        var halfInning = "B ";
        if (game.status.top_inning == "Y") {
          halfInning = "T ";
        }
        displayString = halfInning + game.status.inning;
      }
      return displayString;
    };
  });

warningTrackApp
  .filter('svgIconStrikesOutsHref', function($sce) {
    return function(soId) {
      return $sce.trustAsResourceUrl('strikes-outs.svg#so' + soId);
    };
  });

warningTrackApp
  .filter('svgIconBallsHref', function($sce) {
    return function(ballsId) {
      return $sce.trustAsResourceUrl('balls.svg#b' + ballsId);
    };
  });

warningTrackApp
  .filter('svgIconTeamsHref', function($sce) {
    return function(teamAbbr) {
      return $sce.trustAsResourceUrl('teams.svg#' + teamAbbr);
    };
  });

warningTrackApp
  .filter('leverageToSvg', function($sce) {
    return function(leverageIndex) {
      if (leverageIndex < 1.0) {
        warningId = "0"
      } else
      if (leverageIndex < 1.5) {
        warningId = "1"
      } else
      if (leverageIndex < 2.0) {
        warningId = "2"
      } else
      if (leverageIndex < 2.5) {
        warningId = "3"
      } else
      if (leverageIndex < 3.0) {
        warningId = "4"
      } else {
        warningId = "5"
      }
      return $sce.trustAsResourceUrl('warning.svg#w' + warningId);
    };
  });

warningTrackApp
  .filter('logoPosition', function($filter) {
    return function(id) {
      var yPos = -64 * id - 2;
      return "{'background-position':'-64px " + yPos.toString() + "px'}";
    };
  });

warningTrackApp.controller('WarningTrackCtrl', ['$scope', '$http', '$filter', '$interval',
  function($scope, $http, $filter, $interval) {
    $http.get('/games').success(function(data) {
      $scope.games = data;
    });
    $interval(function() {
      $http.get('/games').success(function(data) {
        $scope.games = data;
      });
    }, 30000);

    $scope.orderProp = ['-leverage_index', 'status.status'];
  }
]);
