#ifndef RELAY8_H_
#define RELAY8_H_

#include <stdint.h>

#define RETRY_TIMES	10
#define RELAY8_INPORT_REG_ADD	0x00
#define RELAY8_OUTPORT_REG_ADD	0x01
#define RELAY8_POLINV_REG_ADD	0x02
#define RELAY8_CFG_REG_ADD		0x03

#define CHANNEL_NR_MIN		1
#define RELAY_CH_NR_MAX		8

#define ERROR	-1
#define OK		0
#define FAIL	-1

#define BOARD_INIT_FAILED 101
#define RELAY_VALUE_OUT_OF_RANGE 102
#define RELAY_STATE_OUT_OF_RANGE 103
#define ERROR_WRITING_RELAY 104
#define ERROR_READING_RELAY 105
#define FAILED_WRITING_RELAY_AFTER_RETRY 106
#define SUCCESS 0

#define RELAY8_HW_I2C_BASE_ADD	0x38
#define RELAY8_HW_I2C_ALTERNATE_BASE_ADD 0x20
typedef uint8_t u8;
typedef uint16_t u16;

typedef enum
{
	OFF = 0,
	ON,
	STATE_COUNT
} OutStateEnumType;


int set_relay(int board, int relay, int status);

int get_relay(int board, int relay, int *status);

#endif
