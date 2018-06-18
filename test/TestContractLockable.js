import expectThrow from "../node_modules/openzeppelin-solidity/test/helpers/expectThrow";
import expectEvent from "../node_modules/openzeppelin-solidity/test/helpers/expectEvent";
import assertRevert from "../node_modules/openzeppelin-solidity/test/helpers/assertRevert";

const SmartDegreeMock = artifacts.require('SmartDegree');

require('chai')
  .use(require('chai-as-promised'))
.should();

contract('SmartDegree', function (accounts) {
  
  let mock;
  
  const [
    owner,
    anyone,
  ] = accounts;

  before(async function () {
    mock = await SmartDegreeMock.new();
  });

  context('in normal conditions', () => {
    it('should not be locked with a valid date', async function () {
          await expectEvent.inTransaction(
            mock.deliverDegree("0x000000000000000000000000000000000000000000000000000000000000001",
                               "0xe7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c",
                               { from: owner }),
            'LogNewSealRecorded'
          );
    });
    it('should be locked with a past date (date -> year 2009)', async function () {
      await expectEvent.inTransaction(
        mock.setDateLimit(991084188, { from: owner }),
        'LogDateLimitUpdated'
      );
      await assertRevert(
        mock.deliverDegree("0x000000000000000000000000000000000000000000000000000000000000002",
                           "0xe7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c",
                           { from: owner }),
      );
    });
  });
});