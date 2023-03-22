<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


class LedgerStorage
{
	#[\JMS\Serializer\Annotation\SerializedName('driver')]
    #[\JMS\Serializer\Annotation\Type('string')]
    public string $driver;
    
    /**
     * $ledgers
     * 
     * @var array<string> $ledgers
     */
	#[\JMS\Serializer\Annotation\SerializedName('ledgers')]
    #[\JMS\Serializer\Annotation\Type('array<string>')]
    public array $ledgers;
    
	public function __construct()
	{
		$this->driver = "";
		$this->ledgers = [];
	}
}
