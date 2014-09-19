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

  // dragging starts here
  var $dragging = null;
  var $ptrX = 0;
  var $ptrY = 0;
  var $origY = 0;
  $.each($('.meal-card'), function(i, obj) {
    $(obj).addClass('meal-card__trans');

    $(obj).mousedown(function(e) {
      $dragging = $(e.target);
      $dragging.removeClass('meal-card__trans');
      $ptrX = e.offsetX;
      $ptrY = e.offsetY;
      $origY = $dragging.offset().top;
      $dragging.addClass('dropzone');
    });

    $(obj).mouseup(function (e) {
      $dragging.addClass('meal-card__trans');
      $dragging.offset({top: $origY});
      $dragging.removeClass('dropzone');
      $dragging = null;
      $ptrX = 0;
      $ptrY = 0;
      $origY = 0;
    });

    // when entering another day, swap the day class
    $(obj).mouseenter(function(e) {
      if ($dragging && $dragging.get(0) !== e.target ) {
        $origY = $(e.target).offset().top;
        swapDays($dragging.get(0), e.target);
      }
    });

  });

  $(document.body).on("mousemove", function(e) {
    if ($dragging) {
      $dragging.offset({ top: (e.pageY-$ptrY) });
    }
  });

}());

