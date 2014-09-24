(function() {
  'use strict';
  /*global $*/

  function openDrawer() {
    if ($('#drawer').hasClass('drawer--open')) {
      return;
    }
    toggleDrawer();
  }

  function closeDrawer() {
    if (!$('#drawer').hasClass('drawer--open')) {
      return;
    }
    toggleDrawer();
  }

  function toggleDrawer() {
    $('#drawer')
      .addClass('trans')
      .toggleClass('drawer--open')
      .on('transitionend', function(e) {
        if (e.target.id === 'drawer') {
          $('#drawer')
            .removeClass('trans')
            .off('transitionend');
        }
      });
  }

  $('#whatis').mousedown(function() {
    openDrawer();
  });

  $('#config').mousedown(function() {
    openDrawer();
  });

  $('#home').mousedown(function() {
    closeDrawer();
  });

  function swapDays(srcObj, dstObj) {
    var srcClass = 'day-' + dayNr(srcObj);
    var dstClass = 'day-' + dayNr(dstObj);
    $(srcObj).removeClass(srcClass).addClass(dstClass);
    $(dstObj).removeClass(dstClass).addClass(srcClass);
  }

  function dayNr(obj) {
    var classes = obj.classList;
    for (var i=0; i<classes.length; i++) {
      var c = classes[i].split('-');
      if (c[0] === 'day') {
        return parseInt(c[1]);
      }
    }
    return -1;
  }

  function lift($elem) {
    $elem.addClass('lift');
    var shadow = $elem.children('paper-shadow');
    shadow.attr('z', "2");
  }

  function drop($elem) {
    $elem.removeClass('lift');
    var shadow = $elem.children('paper-shadow');
    shadow.attr('z', "1");
  }

  function getCurrentBase($elem) {
    var dayNr = $elem.attr('day');
    var base = null;
    $bases.forEach(function(b) {
      if (b.attr('day') === dayNr) { base = b; }
    });
    return base;
  }

  function getBaseAtDay(day) {
    return $('.meal-card__base[day="'+day+'"]');
  }

  function getCardAtDay(day) {
    return $('.meal-card[day="'+day+'"]');
  }

  function resetToDay($elem, day) {
    var base = getBaseAtDay(day);
    $elem.offset(base.offset());
    $elem.attr('day', day);
  }

  function checkBase(top, height) {
    $bases.forEach(function(b) {
      var btop = b.position().top;
      if (b !== $dragging.base &&
          top + height/2 > btop &&
          top < btop + b.height()/2) {
        $dragging.base.removeClass('dropzone');
        var srcDay = $dragging.attr('day');
        var dstDay = b.attr('day');

        $dragging.base = b;
        $dragging.base.addClass('dropzone');
        var card = getCardAtDay(dstDay);
        if (card.length == 1) {
          resetToDay(card, srcDay);
          $dragging.attr('day', dstDay);
        }
        return;
      }
    });
  }

  // dragging starts here
  var $dragging = null;
  var $bases = [];
  var $ptrY = 0;

  //////////////////////
  // Document ready: start setting up stuff
  $(document).ready(function() {

    $.each($('.meal-card__base'), function(i, base) {
      $bases.push($(base));
    });

    $.each($('.meal-card'), function(i, obj) {
      $(obj).addClass('meal-card__trans');

      $(obj).mousedown(function(e) {
        $dragging = $(e.target);
        $dragging.base = getCurrentBase($dragging);
        lift($dragging);
        $dragging.removeClass('meal-card__trans');
        $ptrY = e.offsetY;
      });

      $(obj).mouseup(function (e) {
        drop($dragging);
        $dragging.addClass('meal-card__trans');
        // $dragging.offset($dragging.base.offset());
        // $dragging.attr('day', $dragging.base.attr('day'));
        resetToDay($dragging, $dragging.base.attr('day'));
        $dragging.base.removeClass('dropzone');
        $dragging = null;
        $ptrY = 0;
      });

    });

    $(document.body).on("mousemove", function(e) {
      if ($dragging) {
        $dragging.offset({ top: (e.pageY-$ptrY) });
        checkBase($dragging.position().top, $dragging.height());
      }
    });

  });

}());

