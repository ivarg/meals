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

  $('.drawer__side > paper-fab').mousedown(function() {
    toggleDrawer();
  });

  $('#whatis').mousedown(function() {
    openDrawer();
  });

  $('#config').mousedown(function() {
    openDrawer();
  });

  $('#home').mousedown(function() {
    closeDrawer();
  });

  $('#btn-swap').on('click', function(e) {
    var m = $('.day-1');
    var t = $('.day-2');
    m.removeClass('day-1');
    m.addClass('day-2');
    t.removeClass('day-2');
    t.addClass('day-1');
    console.log('switching days');
  });

  $.each($('.meal-card'), function(i, o) {
    $(o).attr('draggable', 'true');
  });

  function shiftUp(dayNr) {
    var classNow = 'day-' + dayNr;
    var classThen = 'day-' + (dayNr+1);
  }

  function getNrFromClass(className) {
    
  }

  // drag-and-drop event handlers
  $.each($('.meal-card'), function(i, obj) {
    obj.addEventListener('dragstart', function(e) {
      var classes = e.target.classList;
      for (var i=0; i<classes.length; i++) {
        var c = classes[i].split('-');
        if (c[0] === 'day') {
          console.log('storing class ' + classes[i]);
          sessionStorage.setItem('dragDay', classes[i]);
        }
      }
      e.dataTransfer.effectAllowed = 'move';
    });

    obj.addEventListener('dragend', function(e) {
      // e.preventDefault();
    });

    obj.addEventListener('drop', function(e) {
      var d = e.dataTransfer.getData('obj');
      e.preventDefault();
      // debugger;
    });

    obj.addEventListener('dragover', function(e) {
      e.dataTransfer.dropEffect = 'move';
      e.preventDefault();
    });

    // obj.addEventListener('drag', function(e) {
    // });

    // At dragenter, shift the meals
    obj.addEventListener('dragenter', function(e) {
      e.dataTransfer.dropEffect = 'move';
      if (e.target === e.currentTarget) {
        // $(e.target).addClass('dropzone');
        var classes = e.target.classList;
        for (var i=0; i<classes.length; i++) {
          var c = classes[i].split('-');
          var dragDay = sessionStorage.getItem('dragDay');
          if (c[0] === 'day' && classes[i] !== dragDay) {
            console.log('> '+classes[i]);
            swapDays(dragDay, classes[i]);
            break;
          }
        }
      }

      e.preventDefault();
    });

    obj.addEventListener('dragleave', function(e) {
      // if (e.target === e.currentTarget) {
      //   $(e.target).removeClass('dropzone');
      // }
    });

  });

  function swapDays(srcClass, dstClass) {
    console.log('swap days: ' + srcClass + ' <> ' + dstClass);
    var src = $('.'+srcClass);
    var dst = $('.'+dstClass);
    // debugger;
    src.removeClass(srcClass).addClass(dstClass);
    dst.removeClass(dstClass).addClass(srcClass);
    sessionStorage.setItem('dragDay', dstClass);
  }

}());

