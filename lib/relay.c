#include <stdio.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

#include "relay.h"
#include "comm.h"

const u8 relayMaskRemap[8] =
{
	0x01,
	0x04,
	0x40,
	0x10,
	0x20,
	0x80,
	0x08,
	0x02};
const int relayChRemap[8] =
{
	0,
	2,
	6,
	4,
	5,
	7,
	3,
	1};

int relayChSet(int dev, u8 channel, OutStateEnumType state);
int relayChGet(int dev, u8 channel, OutStateEnumType* state);

int doBoardInit(int stack)
{
#ifdef DEBUG
    printf("relay.c: doBoardInit started. stack= %d\r\n", stack);
#endif

	int dev = 0;
	int add = 0;
	uint8_t buff[8];

	if ( (stack < 0) || (stack > 7))
	{
#ifdef DEBUG
		printf("relay.c: doBoardInit invalid stack level. level= %d. valid rage: 0-7\r\n", stack);
#endif
		return ERROR;
	}

/*
	add = (stack + RELAY8_HW_I2C_BASE_ADD) ^ 0x07;
	dev = i2cSetup(add);
	if (dev == -1)
	{
		printf("relay.c: doBoardInit RELAY8_HW_I2C_BASE_ADD i2cSetup error\r\n");
		return ERROR;

	}
*/

	static int stack_file[8] = {-1, -1, -1, -1, -1, -1, -1, -1};

    if (stack_file[stack] == -1) {
        add = (stack + RELAY8_HW_I2C_ALTERNATE_BASE_ADD) ^ 0x07;

        stack_file[stack] = i2cSetup(add);

        if (stack_file[stack] == -1)
        {
#ifdef DEBUG
            printf("relay.c: doBoardInit RELAY8_HW_I2C_ALTERNATE_BASE_ADD i2cSetup error\r\n");
#endif
            return ERROR;
        }
    }

    dev = stack_file[stack];

    if (ERROR == i2cMem8Read(dev, RELAY8_CFG_REG_ADD, buff, 1))
    {
#ifdef DEBUG
        printf("relay.c: doBoardInit i2cMem8Read error\r\n");
#endif
        return ERROR;
    }

	if (buff[0] != 0) //non initialized I/O Expander
	{
		// make all I/O pins output
		buff[0] = 0;
		if (0 > i2cMem8Write(dev, RELAY8_CFG_REG_ADD, buff, 1))
		{
#ifdef DEBUG
		    printf("relay.c: doBoardInit i2cMem8Write error\r\n");
#endif
			return ERROR;
		}
		// put all pins in 0-logic state
		buff[0] = 0;
		if (0 > i2cMem8Write(dev, RELAY8_OUTPORT_REG_ADD, buff, 1))
		{
#ifdef DEBUG
		    printf("relay.c: doBoardInit i2cMem8Write error\r\n");
#endif
			return ERROR;
		}
	}

#ifdef DEBUG
    printf("relay.c: doBoardInit success\r\n");
#endif
	return dev;
}

int relayChSet(int dev, u8 channel, OutStateEnumType state)
{

#ifdef DEBUG
    printf("relay.c: relayChSet started. dev= %d, channel= %d, state= %d\r\n", dev, channel, state);
#endif

	int resp;
	u8 buff[2];

	if ( (channel < CHANNEL_NR_MIN) || (channel > RELAY_CH_NR_MAX))
	{
#ifdef DEBUG
		printf("relay.c: relayChSet invalid relay nr\r\n");
#endif
		return ERROR;
	}
	if (FAIL == i2cMem8Read(dev, RELAY8_INPORT_REG_ADD, buff, 1))
	{
#ifdef DEBUG
		printf("relay.c: relayChSet i2cMem8Read error\r\n");
#endif
		return FAIL;
	}

	switch (state)
	{
	case OFF:
		buff[0] &= ~ (1 << relayChRemap[channel - 1]);
		resp = i2cMem8Write(dev, RELAY8_OUTPORT_REG_ADD, buff, 1);
		break;
	case ON:
		buff[0] |= 1 << relayChRemap[channel - 1];
		resp = i2cMem8Write(dev, RELAY8_OUTPORT_REG_ADD, buff, 1);
		break;
	default:
#ifdef DEBUG
		printf("relay.c: relayChSet unknown relay state\r\n");
#endif
		return ERROR;
		break;
	}

#ifdef DEBUG
	printf("relay.c: relayChSet success\r\n");
#endif
	return resp;
}

int relayChGet(int dev, u8 channel, OutStateEnumType* state)
{
#ifdef DEBUG
    printf("relay.c: relayChGet started. dev= %d, channel= %d, state= %d\r\n", dev, channel, state);
#endif

	u8 buff[2];

	if (NULL == state)
	{
		return ERROR;
	}

	if ( (channel < CHANNEL_NR_MIN) || (channel > RELAY_CH_NR_MAX))
	{
#ifdef DEBUG
		printf("relay.c: relayChGet invalid relay nr.\r\n");
#endif
		return ERROR;
	}

	if (FAIL == i2cMem8Read(dev, RELAY8_INPORT_REG_ADD, buff, 1))
	{
#ifdef DEBUG
	    printf("relay.c: relayChGet i2cMem8Read error.\r\n");
#endif
		return ERROR;
	}

	if (buff[0] & (1 << relayChRemap[channel - 1]))
	{
		*state = ON;
	}
	else
	{
		*state = OFF;
	}

#ifdef DEBUG
	printf("relay.c: relayChGet success\r\n");
#endif
	return OK;
}

int set_relay(int board, int relay, int status) {
#ifdef DEBUG
    printf("relay.c: set_relay board= %d relay= %d\r\n", board, relay);
#endif

	int pin = 0;
	OutStateEnumType state = STATE_COUNT;
	int dev = 0;
	OutStateEnumType stateR = STATE_COUNT;

	int retry = RETRY_TIMES;

	dev = doBoardInit(board);

    while ( (dev <= 0) && (retry > 0) ) {
#ifdef DEBUG
        printf("relay.c: set_relay doBoardInit(board= %d). retrying countdown: %d\r\n", board, retry);
#endif

        dev = doBoardInit(board);
        retry--;
    }

    if (dev <= 0)
    {
#ifdef DEBUG
        printf("relay.c: set_relay doBoardInit(board= %d). no more retries. returning BOARD_INIT_FAILED.\r\n", board);
#endif

        return BOARD_INIT_FAILED;
    }

    pin = relay;

    if ( (pin < CHANNEL_NR_MIN) || (pin > RELAY_CH_NR_MAX))
    {
#ifdef DEBUG
        printf("relay.c: set_relay RELAY_VALUE_OUT_OF_RANGE\r\n");
#endif
 		return RELAY_VALUE_OUT_OF_RANGE;
    }

    if (status == 1) {
        state = ON;
    } else if ( status == 0) {
        state = OFF;
    } else {
#ifdef DEBUG
        printf("relay.c: set_relay RELAY_STATE_OUT_OF_RANGE\r\n");
#endif
        return RELAY_STATE_OUT_OF_RANGE;
    }

    retry = RETRY_TIMES;

    while ((retry > 0) && (stateR != state)) {
        while((retry > 0) && OK != relayChSet(dev, pin, state))
        {
#ifdef DEBUG
            printf("relay.c: set_relay error relayChSet. retry countdown= %d\r\n", retry);
#endif
            retry--;
        }

        while ((retry > 0) && OK != relayChGet(dev, pin, &stateR))
        {
#ifdef DEBUG
            printf("relay.c: set_relay error relayChGet. retry countdown= %d\r\n", retry);
#endif
            retry--;
        }
        retry--;
    }

    if (retry == 0) {
#ifdef DEBUG
        printf("relay.c: set_relay FAILED_WRITING_RELAY_AFTER_RETRY\r\n");
#endif
        return FAILED_WRITING_RELAY_AFTER_RETRY;
    }

#ifdef DEBUG
    printf("relay.c: set_relay success\r\n");
#endif
    return SUCCESS;
}

int get_relay(int board, int relay, int *out_status) {

#ifdef DEBUG
    printf("relay.c: get_relay board= %d relay= %d\r\n", board, relay);
#endif

    int pin = 0;
    int dev = 0;
    OutStateEnumType state = STATE_COUNT;

	int retry = RETRY_TIMES;
	dev = doBoardInit(board);

    while ( (dev <= 0) && (retry > 0) ) {
#ifdef DEBUG
        printf("relay.c: get_relay failed doBoardInit(board= %d). retrying countdown: %d\r\n", board, retry);
#endif

        dev = doBoardInit(board);
        retry--;
    }

    if (dev <= 0)
    {
#ifdef DEBUG
        printf("relay.c: get_relay failed doBoardInit(board= %d). no more retries. returning BOARD_INIT_FAILED.\r\n", board);
#endif

        return BOARD_INIT_FAILED;
    }

    pin = relay;
    if ( (pin < CHANNEL_NR_MIN) || (pin > RELAY_CH_NR_MAX))
    {
#ifdef DEBUG
        printf("relay.c: get_relay RELAY_VALUE_OUT_OF_RANGE\r\n");
#endif
        return RELAY_VALUE_OUT_OF_RANGE;
    }

	retry = RETRY_TIMES;

    while ((retry > 0) && OK != relayChGet(dev, pin, &state))
    {
#ifdef DEBUG
        printf("relay.c: get_relay error relayChGet. retry countdown= %d\r\n", retry);
#endif
        retry--;
    }

    if (retry <= 0)
    {
#ifdef DEBUG
        printf("relay.c: get_relay relayChGet failed. returning ERROR_READING_RELAY\r\n");
#endif
           return ERROR_READING_RELAY;
    }

    if (state != 0)
    {
          *out_status = 1;
    }
    else
    {
        *out_status = 0;
    }

#ifdef DEBUG
    printf("relay.c: get_relay success. out_status= %d\r\n", out_status);
#endif
    return 0;
}
