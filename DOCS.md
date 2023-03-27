# Documentation

## Testdata
The expected result of the testdata (besides not crashing the program) is defined
in a Microsoft Excel sheet. The values in said sheet map to the data structures
in the decoder as follows:

### Profile
The column **Profile** refers to the **seq_profile*. There are 4 different profiles

### Subsampling X/Y
Subsampling information is stored in the color config inside of the sequence header.

### Max picture width/height
The maximum picture width is equivalent to **max_frame_width_minus_1**.
The maximum picture height is equivalent to **max_frame_height_minus_1**.
Note that the variables refer to frames, not pictures. 
There is no reference to a picture width/height in the spec.

### Number of spatial layers

The number of spatial layers is dependent on the syntax element **operating_points_cnt_minus_1**.
It defines how many spatial and temporal layers should be decoded.
**operating_point_idc[i]** indicates which spatial and temporal layers should be decoded for a 
given operating point **i**. If the entry is equal to **0**, then no scalability information
is included. Adding all spatial layers *should* result in the number of spatial layers.

