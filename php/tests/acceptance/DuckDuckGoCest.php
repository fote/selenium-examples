<?php
class DuckDuckGoCest
{
    public function rootWorks(AcceptanceTester $I)
    {
        sleep(5);
        $I->amOnPage('/');
        $I->makeScreenshot('screenshot');
    }
}
