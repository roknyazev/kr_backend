# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.20

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:

#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:

# Disable VCS-based implicit rules.
% : %,v

# Disable VCS-based implicit rules.
% : RCS/%

# Disable VCS-based implicit rules.
% : RCS/%,v

# Disable VCS-based implicit rules.
% : SCCS/s.%

# Disable VCS-based implicit rules.
% : s.%

.SUFFIXES: .hpux_make_needs_suffix_list

# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:
.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /snap/clion/163/bin/cmake/linux/bin/cmake

# The command to remove a file.
RM = /snap/clion/163/bin/cmake/linux/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/romak/Desktop/backend/router

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/romak/Desktop/backend/router/cmake-build-debug

# Include any dependencies generated for this target.
include deps/lemon/lemon/CMakeFiles/lemon.dir/depend.make
# Include the progress variables for this target.
include deps/lemon/lemon/CMakeFiles/lemon.dir/progress.make

# Include the compile flags for this target's objects.
include deps/lemon/lemon/CMakeFiles/lemon.dir/flags.make

deps/lemon/lemon/CMakeFiles/lemon.dir/arg_parser.cc.o: deps/lemon/lemon/CMakeFiles/lemon.dir/flags.make
deps/lemon/lemon/CMakeFiles/lemon.dir/arg_parser.cc.o: ../lemon/lemon/arg_parser.cc
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/romak/Desktop/backend/router/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object deps/lemon/lemon/CMakeFiles/lemon.dir/arg_parser.cc.o"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/lemon.dir/arg_parser.cc.o -c /home/romak/Desktop/backend/router/lemon/lemon/arg_parser.cc

deps/lemon/lemon/CMakeFiles/lemon.dir/arg_parser.cc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/lemon.dir/arg_parser.cc.i"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/romak/Desktop/backend/router/lemon/lemon/arg_parser.cc > CMakeFiles/lemon.dir/arg_parser.cc.i

deps/lemon/lemon/CMakeFiles/lemon.dir/arg_parser.cc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/lemon.dir/arg_parser.cc.s"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/romak/Desktop/backend/router/lemon/lemon/arg_parser.cc -o CMakeFiles/lemon.dir/arg_parser.cc.s

deps/lemon/lemon/CMakeFiles/lemon.dir/base.cc.o: deps/lemon/lemon/CMakeFiles/lemon.dir/flags.make
deps/lemon/lemon/CMakeFiles/lemon.dir/base.cc.o: ../lemon/lemon/base.cc
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/romak/Desktop/backend/router/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building CXX object deps/lemon/lemon/CMakeFiles/lemon.dir/base.cc.o"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/lemon.dir/base.cc.o -c /home/romak/Desktop/backend/router/lemon/lemon/base.cc

deps/lemon/lemon/CMakeFiles/lemon.dir/base.cc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/lemon.dir/base.cc.i"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/romak/Desktop/backend/router/lemon/lemon/base.cc > CMakeFiles/lemon.dir/base.cc.i

deps/lemon/lemon/CMakeFiles/lemon.dir/base.cc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/lemon.dir/base.cc.s"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/romak/Desktop/backend/router/lemon/lemon/base.cc -o CMakeFiles/lemon.dir/base.cc.s

deps/lemon/lemon/CMakeFiles/lemon.dir/color.cc.o: deps/lemon/lemon/CMakeFiles/lemon.dir/flags.make
deps/lemon/lemon/CMakeFiles/lemon.dir/color.cc.o: ../lemon/lemon/color.cc
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/romak/Desktop/backend/router/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Building CXX object deps/lemon/lemon/CMakeFiles/lemon.dir/color.cc.o"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/lemon.dir/color.cc.o -c /home/romak/Desktop/backend/router/lemon/lemon/color.cc

deps/lemon/lemon/CMakeFiles/lemon.dir/color.cc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/lemon.dir/color.cc.i"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/romak/Desktop/backend/router/lemon/lemon/color.cc > CMakeFiles/lemon.dir/color.cc.i

deps/lemon/lemon/CMakeFiles/lemon.dir/color.cc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/lemon.dir/color.cc.s"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/romak/Desktop/backend/router/lemon/lemon/color.cc -o CMakeFiles/lemon.dir/color.cc.s

deps/lemon/lemon/CMakeFiles/lemon.dir/lp_base.cc.o: deps/lemon/lemon/CMakeFiles/lemon.dir/flags.make
deps/lemon/lemon/CMakeFiles/lemon.dir/lp_base.cc.o: ../lemon/lemon/lp_base.cc
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/romak/Desktop/backend/router/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_4) "Building CXX object deps/lemon/lemon/CMakeFiles/lemon.dir/lp_base.cc.o"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/lemon.dir/lp_base.cc.o -c /home/romak/Desktop/backend/router/lemon/lemon/lp_base.cc

deps/lemon/lemon/CMakeFiles/lemon.dir/lp_base.cc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/lemon.dir/lp_base.cc.i"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/romak/Desktop/backend/router/lemon/lemon/lp_base.cc > CMakeFiles/lemon.dir/lp_base.cc.i

deps/lemon/lemon/CMakeFiles/lemon.dir/lp_base.cc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/lemon.dir/lp_base.cc.s"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/romak/Desktop/backend/router/lemon/lemon/lp_base.cc -o CMakeFiles/lemon.dir/lp_base.cc.s

deps/lemon/lemon/CMakeFiles/lemon.dir/lp_skeleton.cc.o: deps/lemon/lemon/CMakeFiles/lemon.dir/flags.make
deps/lemon/lemon/CMakeFiles/lemon.dir/lp_skeleton.cc.o: ../lemon/lemon/lp_skeleton.cc
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/romak/Desktop/backend/router/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_5) "Building CXX object deps/lemon/lemon/CMakeFiles/lemon.dir/lp_skeleton.cc.o"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/lemon.dir/lp_skeleton.cc.o -c /home/romak/Desktop/backend/router/lemon/lemon/lp_skeleton.cc

deps/lemon/lemon/CMakeFiles/lemon.dir/lp_skeleton.cc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/lemon.dir/lp_skeleton.cc.i"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/romak/Desktop/backend/router/lemon/lemon/lp_skeleton.cc > CMakeFiles/lemon.dir/lp_skeleton.cc.i

deps/lemon/lemon/CMakeFiles/lemon.dir/lp_skeleton.cc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/lemon.dir/lp_skeleton.cc.s"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/romak/Desktop/backend/router/lemon/lemon/lp_skeleton.cc -o CMakeFiles/lemon.dir/lp_skeleton.cc.s

deps/lemon/lemon/CMakeFiles/lemon.dir/random.cc.o: deps/lemon/lemon/CMakeFiles/lemon.dir/flags.make
deps/lemon/lemon/CMakeFiles/lemon.dir/random.cc.o: ../lemon/lemon/random.cc
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/romak/Desktop/backend/router/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_6) "Building CXX object deps/lemon/lemon/CMakeFiles/lemon.dir/random.cc.o"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/lemon.dir/random.cc.o -c /home/romak/Desktop/backend/router/lemon/lemon/random.cc

deps/lemon/lemon/CMakeFiles/lemon.dir/random.cc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/lemon.dir/random.cc.i"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/romak/Desktop/backend/router/lemon/lemon/random.cc > CMakeFiles/lemon.dir/random.cc.i

deps/lemon/lemon/CMakeFiles/lemon.dir/random.cc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/lemon.dir/random.cc.s"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/romak/Desktop/backend/router/lemon/lemon/random.cc -o CMakeFiles/lemon.dir/random.cc.s

deps/lemon/lemon/CMakeFiles/lemon.dir/bits/windows.cc.o: deps/lemon/lemon/CMakeFiles/lemon.dir/flags.make
deps/lemon/lemon/CMakeFiles/lemon.dir/bits/windows.cc.o: ../lemon/lemon/bits/windows.cc
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/romak/Desktop/backend/router/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_7) "Building CXX object deps/lemon/lemon/CMakeFiles/lemon.dir/bits/windows.cc.o"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/lemon.dir/bits/windows.cc.o -c /home/romak/Desktop/backend/router/lemon/lemon/bits/windows.cc

deps/lemon/lemon/CMakeFiles/lemon.dir/bits/windows.cc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/lemon.dir/bits/windows.cc.i"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/romak/Desktop/backend/router/lemon/lemon/bits/windows.cc > CMakeFiles/lemon.dir/bits/windows.cc.i

deps/lemon/lemon/CMakeFiles/lemon.dir/bits/windows.cc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/lemon.dir/bits/windows.cc.s"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/romak/Desktop/backend/router/lemon/lemon/bits/windows.cc -o CMakeFiles/lemon.dir/bits/windows.cc.s

# Object files for target lemon
lemon_OBJECTS = \
"CMakeFiles/lemon.dir/arg_parser.cc.o" \
"CMakeFiles/lemon.dir/base.cc.o" \
"CMakeFiles/lemon.dir/color.cc.o" \
"CMakeFiles/lemon.dir/lp_base.cc.o" \
"CMakeFiles/lemon.dir/lp_skeleton.cc.o" \
"CMakeFiles/lemon.dir/random.cc.o" \
"CMakeFiles/lemon.dir/bits/windows.cc.o"

# External object files for target lemon
lemon_EXTERNAL_OBJECTS =

deps/lemon/lemon/libemon.a: deps/lemon/lemon/CMakeFiles/lemon.dir/arg_parser.cc.o
deps/lemon/lemon/libemon.a: deps/lemon/lemon/CMakeFiles/lemon.dir/base.cc.o
deps/lemon/lemon/libemon.a: deps/lemon/lemon/CMakeFiles/lemon.dir/color.cc.o
deps/lemon/lemon/libemon.a: deps/lemon/lemon/CMakeFiles/lemon.dir/lp_base.cc.o
deps/lemon/lemon/libemon.a: deps/lemon/lemon/CMakeFiles/lemon.dir/lp_skeleton.cc.o
deps/lemon/lemon/libemon.a: deps/lemon/lemon/CMakeFiles/lemon.dir/random.cc.o
deps/lemon/lemon/libemon.a: deps/lemon/lemon/CMakeFiles/lemon.dir/bits/windows.cc.o
deps/lemon/lemon/libemon.a: deps/lemon/lemon/CMakeFiles/lemon.dir/build.make
deps/lemon/lemon/libemon.a: deps/lemon/lemon/CMakeFiles/lemon.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/romak/Desktop/backend/router/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_8) "Linking CXX static library libemon.a"
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && $(CMAKE_COMMAND) -P CMakeFiles/lemon.dir/cmake_clean_target.cmake
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && $(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/lemon.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
deps/lemon/lemon/CMakeFiles/lemon.dir/build: deps/lemon/lemon/libemon.a
.PHONY : deps/lemon/lemon/CMakeFiles/lemon.dir/build

deps/lemon/lemon/CMakeFiles/lemon.dir/clean:
	cd /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon && $(CMAKE_COMMAND) -P CMakeFiles/lemon.dir/cmake_clean.cmake
.PHONY : deps/lemon/lemon/CMakeFiles/lemon.dir/clean

deps/lemon/lemon/CMakeFiles/lemon.dir/depend:
	cd /home/romak/Desktop/backend/router/cmake-build-debug && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/romak/Desktop/backend/router /home/romak/Desktop/backend/router/lemon/lemon /home/romak/Desktop/backend/router/cmake-build-debug /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon /home/romak/Desktop/backend/router/cmake-build-debug/deps/lemon/lemon/CMakeFiles/lemon.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : deps/lemon/lemon/CMakeFiles/lemon.dir/depend

