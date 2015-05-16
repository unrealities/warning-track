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
      var displayString = "";

      if (game.status.state < 3) {
        displayString = "Final";
      } else if (game.status.state == 3) {
        displayString = "Postponed";
      } else if (game.status.state > 10 && game.status.state < 20) {
        var d = new Date(game.date_time);
        var ld = new Date(d.getTime()+d.getTimezoneOffset()*60*1000);
        var offset = d.getTimezoneOffset() / 60;
        var hours = d.getHours();
        ld.setHours(hours - offset);

        var ampm = hr < 12 ? "am" : "pm";
        var hr = ld.getHours();
        if (hr > 12) {
          hr = hr-12;
        }
        var min = ld.getMinutes();
        if (min < 10) {
            min = "0" + min;
        }
        displayString = hr + ":" + min + " " + ampm;
      } else if (game.status.state == 21) {
        displayString = "Delayed";
      } else {
        var halfInning = "B ";
        if (game.status.half_inning == "Top") {
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
      var yPos = (-64 * (id-1)) - 1.128*(id-1);
      return "{'background-position':'-64px " + yPos.toString() + "px'}";
    };
  });

warningTrackApp
  .filter('gameDate', function(){
    return function(games) {
      var result = [];
      var today = new Date();
      var today_utc = new Date(today.getUTCFullYear(), today.getUTCMonth(), today.getUTCDate(), today.getUTCHours());
      var view_date = new Date(today.getUTCFullYear(), today.getUTCMonth(), today.getUTCDate());
      if (today_utc.getHours() < 12) {
        view_date = new Date(view_date.getTime() - 1000*60*60*24);
      }
      angular.forEach(games, function(game) {
        var game_date = new Date(game.date_time);
        if(game_date.getDate() == view_date.getDate() &&
           game_date.getMonth() == view_date.getMonth() &&
           game_date.getYear() == view_date.getYear()) {
          result.push(game);
        }
      });
      return result;
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
    }, 80000);

    $scope.orderProp = ['-status.leverage_index', '-status.state', 'date_time'];
  }
]);
