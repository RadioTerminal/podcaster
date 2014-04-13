package utils

/*
#cgo LDFLAGS: -lgroove

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <groove/groove.h>
#include <limits.h>

//compile with gcc -o waveform main.c -O3 -lgroove -std=c99

// Define a vector type
typedef struct {
  int size;
  float *data;
  float duration;
} Samples;

int getSoundPoints(Samples* soundPoints, char* in_file_path, int fps) {

    float min_sample = -1.0f;
    int samples_to_take = 0;
    int frame_count = 0;
    int frames_until_emit;
    int emit_every;

    printf("Using libgroove v%s\n", groove_version());
    groove_init();
    groove_set_logging(GROOVE_LOG_INFO);
    struct GrooveFile *file = groove_file_open(in_file_path);
    if (! file) {
        fprintf(stderr, "Error opening input file: %s\n", in_file_path);
        groove_finish();
        return 1;
    }

    float duration = groove_file_duration(file);

    struct GroovePlaylist *playlist = groove_playlist_create();
    struct GrooveSink *sink = groove_sink_create();
    struct GrooveBuffer *buffer;


    sink->audio_format.sample_rate = 44100;
    sink->audio_format.channel_layout = GROOVE_CH_LAYOUT_MONO;
    sink->audio_format.sample_fmt = GROOVE_SAMPLE_FMT_FLT;
    if (groove_sink_attach(sink, playlist) < 0) {
        fprintf(stderr, "error attaching sink\n");
        return 1;
    }

    struct GroovePlaylistItem *item =
        groove_playlist_insert(playlist, file, 1.0, NULL);

    // scan the song for the exact correct duration
    while (groove_sink_buffer_get(sink, &buffer, 1) == GROOVE_BUFFER_YES) {
        frame_count += buffer->frame_count;
        groove_buffer_unref(buffer);
    }
    groove_playlist_seek(playlist, item, 0);


    // Calculate best sample resolution for file
    samples_to_take = duration * fps; // take sample every 125ms (at 8)
    emit_every = frame_count / samples_to_take;
    frames_until_emit = emit_every;

    // Initialize vector
    soundPoints->size = 0;
    soundPoints->data = calloc(samples_to_take+1, sizeof(float));
    soundPoints->duration = groove_file_duration(file);
    int i;
    while (groove_sink_buffer_get(sink, &buffer, 1) == GROOVE_BUFFER_YES) {
        // process the buffer
        for (i = 0; i < buffer->frame_count && soundPoints->size < samples_to_take;
                i += 1, frames_until_emit -= 1)
        {
            if (frames_until_emit == 0) {
                soundPoints->data[soundPoints->size++] = min_sample;
                frames_until_emit = emit_every;
                min_sample = -1.0f;
            }
            float *samples = (float *) buffer->data[0];
            float sample = samples[i];
            if (sample > min_sample) min_sample = sample;
        }

        groove_buffer_unref(buffer);
    }

    // emit the last column if necessary. This will have to run multiple times
    // if the duration specified in the metadata is incorrect.
    while (soundPoints->size < samples_to_take) {
        soundPoints->data[soundPoints->size++] = 0.0f;
    }

    groove_sink_detach(sink);
    groove_sink_destroy(sink);

    groove_playlist_clear(playlist);
    groove_file_close(file);
    groove_playlist_destroy(playlist);
    groove_finish();
}

float getSample(Samples *vector, int index){
    return vector->data[index];
}

*/
import "C"

import (
	"runtime"
	"strconv"
	"unsafe"
)

type Samples struct {
	size     int
	data     *float64
	duration float64
}

func GenerateSamplesAsFloat(file string) ([]float64, float64) {
	var sample Samples
	runtime.LockOSThread()
	_file := C.CString(file)
	defer C.free(unsafe.Pointer(_file))

	C.getSoundPoints((*_Ctype_Samples)(unsafe.Pointer(&sample)), _file, 4)
	defer C.free(unsafe.Pointer(sample.data))
	var data []float64
	for i := 0; i < sample.size-1; i++ {
		var f float64 = (float64)(C.getSample((*_Ctype_Samples)(unsafe.Pointer(&sample)), C.int(i)))
		data = append(data, f)
	}
	runtime.UnlockOSThread()
	return data, sample.duration
}

func GenerateSamplesAsString(file string, precision int) ([]string, float64) {
	var data []string
	datas, duration := GenerateSamplesAsFloat(file)
	for _, f := range datas {
		data = append(data, strconv.FormatFloat(f, 'f', precision, 32))
	}
	return data, duration
}
