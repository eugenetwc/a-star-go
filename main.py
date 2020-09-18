from cffi import FFI

ffi = FFI()

ffi.cdef("""
    typedef long long GoInt;

    typedef struct {
        void* data;
        GoInt len;
        GoInt cap;
    } GoSlice;

    typedef struct {
        const char *data;
        GoInt len;
    } GoString;

    typedef struct {
        GoInt x;
        GoInt y;
    } Location;

    GoSlice AStar(GoSlice coordMap);
"""
)
lib = ffi.dlopen("./_astar.so")

maze = [
        [0, 0, 1, 0, 3, 0],
		[0, 0, 0, 1, 0, 0],
		[1, 1, 0, 1, 1, 0],
		[2, 0, 0, 0, 0, 0],
    ]

data = ffi.new(f"GoInt[{len(maze)}][{len(maze[0])}]")
for i in range(len(maze)):
    data[i] = ffi.new(f"GoInt[{len(maze[i])}]", maze[i])
nums = ffi.new("GoSlice*", {'data': data, 'len':len(maze)*len(maze[0]), 'cap':len(maze)*len(maze[0])})

test = lib.AStar(nums[0])
print(test)




