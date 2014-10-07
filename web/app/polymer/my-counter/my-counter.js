/*global Polymer*/

(function() {
  'use strict';
  new Polymer({
    counter: 0, // Default value
    counterChanged: function() {
      this.$.counterVal.classList.add('highlight');
    },
    increment: function() {
      this.counter++;
    }
  });
})();


