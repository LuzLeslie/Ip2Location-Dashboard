<template>
  <div id="wrapper">

    <!-- Sidebar -->
    <ul class="navbar-nav bg-gradient-primary sidebar sidebar-dark accordion" >

      <!-- Sidebar - Brand -->
      <a class="sidebar-brand d-flex align-items-center justify-content-center" href="/">
        <div class="sidebar-brand-icon rotate-n-15">
          <i class="fas fa-laugh-wink"></i>
        </div>
        <div class="sidebar-brand-text mx-3">Ip2Location</div>
      </a>

      <!-- Divider -->
      <hr class="sidebar-divider my-0">

      <!-- Nav Item - Dashboard -->
      <li class="nav-item active">
        <a class="nav-link" href="/">
          <font-awesome-icon icon="tachometer-alt" />
          <span> Dashboard</span></a>
      </li>

      <!-- Divider -->
      <hr class="sidebar-divider">

      <!-- Heading -->
      <div class="sidebar-heading">
        Tools
      </div>

      <!-- Nav Item - Pages Collapse Menu -->
      <li class="nav-item">
        <a class="nav-link">
           <font-awesome-icon icon="table" />
          <router-link to="/log" class="routerlink"> Analysis log</router-link>
        </a>
      </li>

      <!-- Nav Item - Utilities Collapse Menu -->
      <li class="nav-item">
        <a class="nav-link collapsed" href="#" >
          <font-awesome-icon icon="search-location" />
          <router-link to="/analyze" class="routerlink"> Analyze</router-link>
        </a>
      </li>

      <!-- Divider -->
      <hr class="sidebar-divider">

      <!-- Heading -->
      <div class="sidebar-heading">
        Addons
      </div>

      <!-- Nav Item - Charts -->
      <li class="nav-item">
        <a class="nav-link" href="#">
          <font-awesome-icon icon="info-circle" />
          <span> Info</span>
        </a>
      </li>

      <!-- Divider -->
      <hr class="sidebar-divider d-none d-md-block">

    </ul>
    <!-- End of Sidebar -->

    <!-- Content Wrapper -->
    <div id="content-wrapper" class="d-flex flex-column">

      <!-- Main Content -->
      <div id="content">

        <!-- Topbar -->
        <nav class="navbar navbar-expand navbar-light bg-white topbar mb-4 static-top shadow">
          <h4>{{tittleBar}}</h4>
        </nav>
        <!-- End of Topbar -->

        <!-- Begin Page Content -->
        <div class="container-fluid">
          <router-view/>
        </div>
        <!-- /.container-fluid -->

      </div>
      <!-- End of Main Content -->

      <!-- Footer -->
      <footer class="sticky-footer bg-white">
        <div class="container my-auto">
          <div class="copyright text-center my-auto">
            <span>ip2location.com</span>
          </div>
        </div>
      </footer>
      <!-- End of Footer -->

    </div>
    <!-- End of Content Wrapper -->

  </div>
</template>

<script>
import { mapState } from 'vuex'
export default {
  name: 'App',
  computed: {
    ...mapState([
      'tittleBar'
    ])
  }
}
</script>

<style>

.routerlink {
  color: #FFF;
}

.routerlink:hover {
   color: white;
   font-weight: bold;
}

#wrapper {
  display: flex;
}

#wrapper #content-wrapper {
  background-color: #f8f9fc;
  width: 100%;
  overflow-x: hidden;
}

#wrapper #content-wrapper #content {
  flex: 1 0 auto;
}

.bg-gradient-primary {
  background-color: #1172D2;
  background-image: linear-gradient(180deg, #1172D2 10%, #224abe 100%);
  background-size: cover;
}

/*---*/

.bg-gradient-secondary {
  background-color: #858796;
  background-image: linear-gradient(180deg, #858796 10%, #60616f 100%);
  background-size: cover;
}

.bg-gradient-success {
  background-color: #1cc88a;
  background-image: linear-gradient(180deg, #1cc88a 10%, #13855c 100%);
  background-size: cover;
}

.bg-gradient-info {
  background-color: #36b9cc;
  background-image: linear-gradient(180deg, #36b9cc 10%, #258391 100%);
  background-size: cover;
}

.bg-gradient-warning {
  background-color: #f6c23e;
  background-image: linear-gradient(180deg, #f6c23e 10%, #dda20a 100%);
  background-size: cover;
}

.bg-gradient-danger {
  background-color: #e74a3b;
  background-image: linear-gradient(180deg, #e74a3b 10%, #be2617 100%);
  background-size: cover;
}

.bg-gradient-light {
  background-color: #f8f9fc;
  background-image: linear-gradient(180deg, #f8f9fc 10%, #c2cbe5 100%);
  background-size: cover;
}

.bg-gradient-dark {
  background-color: #5a5c69;
  background-image: linear-gradient(180deg, #5a5c69 10%, #373840 100%);
  background-size: cover;
}

.bg-gray-100 {
  background-color: #f8f9fc !important;
}

.bg-gray-200 {
  background-color: #eaecf4 !important;
}

.bg-gray-300 {
  background-color: #dddfeb !important;
}

.bg-gray-400 {
  background-color: #d1d3e2 !important;
}

.bg-gray-500 {
  background-color: #b7b9cc !important;
}

.bg-gray-600 {
  background-color: #858796 !important;
}

.bg-gray-700 {
  background-color: #6e707e !important;
}

.bg-gray-800 {
  background-color: #5a5c69 !important;
}

.bg-gray-900 {
  background-color: #3a3b45 !important;
}

.topbar {
  height: 4.375rem;
}

.topbar #sidebarToggleTop {
  height: 2.5rem;
  width: 2.5rem;
}

.topbar #sidebarToggleTop:hover {
  background-color: #eaecf4;
}

.topbar #sidebarToggleTop:active {
  background-color: #dddfeb;
}

.topbar .navbar-search {
  width: 25rem;
}

.topbar .navbar-search input {
  font-size: 0.85rem;
  height: auto;
}

.topbar .topbar-divider {
  width: 0;
  border-right: 1px solid #e3e6f0;
  height: calc(4.375rem - 2rem);
  margin: auto 1rem;
}

.topbar .nav-item .nav-link {
  height: 4.375rem;
  display: flex;
  align-items: center;
  padding: 0 0.75rem;
}

.topbar .nav-item .nav-link:focus {
  outline: none;
}

.topbar .nav-item:focus {
  outline: none;
}

.topbar .dropdown {
  position: static;
}

.topbar .dropdown .dropdown-menu {
  width: calc(100% - 1.5rem);
  right: 0.75rem;
}

.topbar .dropdown-list {
  padding: 0;
  border: none;
  overflow: hidden;
}

.topbar .dropdown-list .dropdown-header {
  background-color: #4e73df;
  border: 1px solid #4e73df;
  padding-top: 0.75rem;
  padding-bottom: 0.75rem;
  color: #fff;
}

.topbar .dropdown-list .dropdown-item {
  white-space: normal;
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
  border-left: 1px solid #e3e6f0;
  border-right: 1px solid #e3e6f0;
  border-bottom: 1px solid #e3e6f0;
  line-height: 1.3rem;
}

.topbar .dropdown-list .dropdown-item .dropdown-list-image {
  position: relative;
  height: 2.5rem;
  width: 2.5rem;
}

.topbar .dropdown-list .dropdown-item .dropdown-list-image img {
  height: 2.5rem;
  width: 2.5rem;
}

.topbar .dropdown-list .dropdown-item .dropdown-list-image .status-indicator {
  background-color: #eaecf4;
  height: 0.75rem;
  width: 0.75rem;
  border-radius: 100%;
  position: absolute;
  bottom: 0;
  right: 0;
  border: 0.125rem solid #fff;
}

.topbar .dropdown-list .dropdown-item .text-truncate {
  max-width: 10rem;
}

.topbar .dropdown-list .dropdown-item:active {
  background-color: #eaecf4;
  color: #3a3b45;
}

.topbar.navbar-dark .navbar-nav .nav-item .nav-link {
  color: rgba(255, 255, 255, 0.8);
}

.topbar.navbar-dark .navbar-nav .nav-item .nav-link:hover {
  color: #fff;
}

.topbar.navbar-dark .navbar-nav .nav-item .nav-link:active {
  color: #fff;
}

.topbar.navbar-light .navbar-nav .nav-item .nav-link {
  color: #d1d3e2;
}

.topbar.navbar-light .navbar-nav .nav-item .nav-link:hover {
  color: #b7b9cc;
}

.topbar.navbar-light .navbar-nav .nav-item .nav-link:active {
  color: #858796;
}

.sidebar {
  width: 6.5rem;
  min-height: 100vh;
}

.sidebar .nav-item {
  position: relative;
}

.sidebar .nav-item:last-child {
  margin-bottom: 1rem;
}

.sidebar .nav-item .nav-link {
  text-align: center;
  padding: 0.75rem 1rem;
  width: 6.5rem;
}

.sidebar .nav-item .nav-link span {
  font-size: 0.65rem;
  display: block;
}

.sidebar .nav-item.active .nav-link {
  font-weight: 700;
}

.sidebar .nav-item .collapse {
  position: absolute;
  left: calc(6.5rem + 1.5rem / 2);
  z-index: 1;
  top: 2px;
}

.sidebar .nav-item .collapse .collapse-inner {
  border-radius: 0.35rem;
  box-shadow: 0 0.15rem 1.75rem 0 rgba(58, 59, 69, 0.15);
}

.sidebar .nav-item .collapsing {
  display: none;
  transition: none;
}

.sidebar .nav-item .collapse .collapse-inner,
.sidebar .nav-item .collapsing .collapse-inner {
  padding: .5rem 0;
  min-width: 10rem;
  font-size: 0.85rem;
  margin: 0 0 1rem 0;
}

.sidebar .nav-item .collapse .collapse-inner .collapse-header,
.sidebar .nav-item .collapsing .collapse-inner .collapse-header {
  margin: 0;
  white-space: nowrap;
  padding: .5rem 1.5rem;
  text-transform: uppercase;
  font-weight: 800;
  font-size: 0.65rem;
  color: #b7b9cc;
}

.sidebar .nav-item .collapse .collapse-inner .collapse-item,
.sidebar .nav-item .collapsing .collapse-inner .collapse-item {
  padding: 0.5rem 1rem;
  margin: 0 0.5rem;
  display: block;
  color: #3a3b45;
  text-decoration: none;
  border-radius: 0.35rem;
  white-space: nowrap;
}

.sidebar .nav-item .collapse .collapse-inner .collapse-item:hover,
.sidebar .nav-item .collapsing .collapse-inner .collapse-item:hover {
  background-color: #eaecf4;
}

.sidebar .nav-item .collapse .collapse-inner .collapse-item:active,
.sidebar .nav-item .collapsing .collapse-inner .collapse-item:active {
  background-color: #dddfeb;
}

.sidebar .nav-item .collapse .collapse-inner .collapse-item.active,
.sidebar .nav-item .collapsing .collapse-inner .collapse-item.active {
  color: #4e73df;
  font-weight: 700;
}

.sidebar #sidebarToggle {
  width: 2.5rem;
  height: 2.5rem;
  text-align: center;
  margin-bottom: 1rem;
  cursor: pointer;
}

.sidebar #sidebarToggle::after {
  font-weight: 900;
  content: '\f104';
  font-family: 'Font Awesome 5 Free';
  margin-right: 0.1rem;
}

.sidebar #sidebarToggle:hover {
  text-decoration: none;
}

.sidebar #sidebarToggle:focus {
  outline: none;
}

.sidebar.toggled {
  width: 0 !important;
  overflow: hidden;
}

.sidebar.toggled #sidebarToggle::after {
  content: '\f105';
  font-family: 'Font Awesome 5 Free';
  margin-left: 0.25rem;
}

.sidebar .sidebar-brand {
  height: 4.375rem;
  text-decoration: none;
  font-size: 1rem;
  font-weight: 800;
  padding: 1.5rem 1rem;
  text-align: center;
  text-transform: uppercase;
  letter-spacing: 0.05rem;
  z-index: 1;
}

.sidebar .sidebar-brand .sidebar-brand-icon i {
  font-size: 2rem;
}

.sidebar .sidebar-brand .sidebar-brand-text {
  display: none;
}

.sidebar hr.sidebar-divider {
  margin: 0 1rem 1rem;
}

.sidebar .sidebar-heading {
  text-align: center;
  padding: 0 1rem;
  font-weight: 800;
  font-size: 0.65rem;
}

@media (min-width: 768px) {
  .sidebar {
    width: 14rem !important;
  }
  .sidebar .nav-item .collapse {
    position: relative;
    left: 0;
    z-index: 1;
    top: 0;
    -webkit-animation: none;
    animation: none;
  }
  .sidebar .nav-item .collapse .collapse-inner {
    border-radius: 0;
    box-shadow: none;
  }
  .sidebar .nav-item .collapsing {
    display: block;
    transition: height 0.15s ease;
  }
  .sidebar .nav-item .collapse,
  .sidebar .nav-item .collapsing {
    margin: 0 1rem;
  }
  .sidebar .nav-item .nav-link {
    display: block;
    width: 100%;
    text-align: left;
    padding: 1rem;
    width: 14rem;
  }
  .sidebar .nav-item .nav-link i {
    font-size: 0.85rem;
    margin-right: 0.25rem;
  }
  .sidebar .nav-item .nav-link span {
    font-size: 0.85rem;
    display: inline;
  }
  .sidebar .nav-item .nav-link[data-toggle="collapse"]::after {
    width: 1rem;
    text-align: center;
    float: right;
    vertical-align: 0;
    border: 0;
    font-weight: 900;
    content: '\f107';
    font-family: 'Font Awesome 5 Free';
  }
  .sidebar .nav-item .nav-link[data-toggle="collapse"].collapsed::after {
    content: '\f105';
  }
  .sidebar .sidebar-brand .sidebar-brand-icon i {
    font-size: 2rem;
  }
  .sidebar .sidebar-brand .sidebar-brand-text {
    display: inline;
  }
  .sidebar .sidebar-heading {
    text-align: left;
  }
  .sidebar.toggled {
    overflow: visible;
    width: 6.5rem !important;
  }
  .sidebar.toggled .nav-item .collapse {
    position: absolute;
    left: calc(6.5rem + 1.5rem / 2);
    z-index: 1;
    top: 2px;
    -webkit-animation-name: growIn;
    animation-name: growIn;
    -webkit-animation-duration: 200ms;
    animation-duration: 200ms;
    -webkit-animation-timing-function: transform cubic-bezier(0.18, 1.25, 0.4, 1), opacity cubic-bezier(0, 1, 0.4, 1);
    animation-timing-function: transform cubic-bezier(0.18, 1.25, 0.4, 1), opacity cubic-bezier(0, 1, 0.4, 1);
  }
  .sidebar.toggled .nav-item .collapse .collapse-inner {
    box-shadow: 0 0.15rem 1.75rem 0 rgba(58, 59, 69, 0.15);
    border-radius: 0.35rem;
  }
  .sidebar.toggled .nav-item .collapsing {
    display: none;
    transition: none;
  }
  .sidebar.toggled .nav-item .collapse,
  .sidebar.toggled .nav-item .collapsing {
    margin: 0;
  }
  .sidebar.toggled .nav-item:last-child {
    margin-bottom: 1rem;
  }
  .sidebar.toggled .nav-item .nav-link {
    text-align: center;
    padding: 0.75rem 1rem;
    width: 6.5rem;
  }
  .sidebar.toggled .nav-item .nav-link span {
    font-size: 0.65rem;
    display: block;
  }
  .sidebar.toggled .nav-item .nav-link i {
    margin-right: 0;
  }
  .sidebar.toggled .nav-item .nav-link[data-toggle="collapse"]::after {
    display: none;
  }
  .sidebar.toggled .sidebar-brand .sidebar-brand-icon i {
    font-size: 2rem;
  }
  .sidebar.toggled .sidebar-brand .sidebar-brand-text {
    display: none;
  }
  .sidebar.toggled .sidebar-heading {
    text-align: center;
  }
}

.sidebar-light .sidebar-brand {
  color: #6e707e;
}

.sidebar-light hr.sidebar-divider {
  border-top: 1px solid #eaecf4;
}

.sidebar-light .sidebar-heading {
  color: #b7b9cc;
}

.sidebar-light .nav-item .nav-link {
  color: #858796;
}

.sidebar-light .nav-item .nav-link i {
  color: #d1d3e2;
}

.sidebar-light .nav-item .nav-link:active, .sidebar-light .nav-item .nav-link:focus, .sidebar-light .nav-item .nav-link:hover {
  color: #6e707e;
}

.sidebar-light .nav-item .nav-link:active i, .sidebar-light .nav-item .nav-link:focus i, .sidebar-light .nav-item .nav-link:hover i {
  color: #6e707e;
}

.sidebar-light .nav-item .nav-link[data-toggle="collapse"]::after {
  color: #b7b9cc;
}

.sidebar-light .nav-item.active .nav-link {
  color: #6e707e;
}

.sidebar-light .nav-item.active .nav-link i {
  color: #6e707e;
}

.sidebar-light #sidebarToggle {
  background-color: #eaecf4;
}

.sidebar-light #sidebarToggle::after {
  color: #b7b9cc;
}

.sidebar-light #sidebarToggle:hover {
  background-color: #dddfeb;
}

.sidebar-dark .sidebar-brand {
  color: #fff;
}

.sidebar-dark hr.sidebar-divider {
  border-top: 1px solid rgba(255, 255, 255, 0.15);
}

.sidebar-dark .sidebar-heading {
  color: rgba(255, 255, 255, 0.4);
}

.sidebar-dark .nav-item .nav-link {
  color: rgba(255, 255, 255, 0.8);
}

.sidebar-dark .nav-item .nav-link i {
  color: rgba(255, 255, 255, 0.3);
}

.sidebar-dark .nav-item .nav-link:active, .sidebar-dark .nav-item .nav-link:focus, .sidebar-dark .nav-item .nav-link:hover {
  color: #fff;
}

.sidebar-dark .nav-item .nav-link:active i, .sidebar-dark .nav-item .nav-link:focus i, .sidebar-dark .nav-item .nav-link:hover i {
  color: #fff;
}

.sidebar-dark .nav-item .nav-link[data-toggle="collapse"]::after {
  color: rgba(255, 255, 255, 0.5);
}

.sidebar-dark .nav-item.active .nav-link {
  color: #fff;
}

.sidebar-dark .nav-item.active .nav-link i {
  color: #fff;
}

.sidebar-dark #sidebarToggle {
  background-color: rgba(255, 255, 255, 0.2);
}

.sidebar-dark #sidebarToggle::after {
  color: rgba(255, 255, 255, 0.5);
}

.sidebar-dark #sidebarToggle:hover {
  background-color: rgba(255, 255, 255, 0.25);
}

.sidebar-dark.toggled #sidebarToggle::after {
  color: rgba(255, 255, 255, 0.5);
}
</style>
